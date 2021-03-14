package main

import (
	"SA-final/user_microservice/models"
	"SA-final/user_microservice/proto/users"
	"SA-final/user_microservice/repository"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
)

type Server struct {
	app *application
	users.UnimplementedUserServiceServer
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo repository.IUserRepository
	conn users.UserServiceClient
}

func (app *application) getBook(c users.UserServiceClient, bookName string) *models.Book {
	ctx := context.Background()
	request := &users.UserAddBookRequest{BookName: bookName}
	response, err := c.AddBook(ctx, request)
	if err != nil {
		log.Fatalf("error while calling AddBook RPC %v", err)
	}
	book := models.Book{
		ID:     int(response.Book.Id),
		Name:   response.Book.Name,
		Author: response.Book.Author,
	}
	return &book
}

func (s *Server) GetAllUsers(ctx context.Context, req *users.GetAllUsersRequest) (*users.GetAllUsersResponse, error) {
	userList, err := s.app.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	protoUsers := []*users.User{}
	for _, user := range userList {
		protoIds := []int32{}
		for _, id := range user.BookIds {
			protoIds = append(protoIds, int32(id))
		}
		protoUser := &users.User{
			Id:       int32(user.ID),
			Email:    user.Email,
			Username: user.Username,
			Password: user.Password,
			BookIds:  protoIds,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	res := &users.GetAllUsersResponse{
		Users: protoUsers,
	}
	return res, err
}

func main()  {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := pgxpool.Connect(context.Background(), "user=postgres password=postgres host=db port=5432 dbname=postgres sslmode=disable pool_max_conns=10")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pool.Close()

	conn, err := grpc.Dial("booksapi:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := users.NewUserServiceClient(conn)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		userRepo: &repository.UserRepository{Pool: pool},
		conn: c,
	}

	server := &http.Server{
		Addr:     ":8002",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	go GRPCServe(app)

	infoLog.Print("Starting server on port 8002")
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func GRPCServe(app *application) {
	l, err := net.Listen("tcp", "usersapi:8003")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	users.RegisterUserServiceServer(s, &Server{app, users.UnimplementedUserServiceServer{}})

	log.Println("Server is running on port:8003")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
