package main

import (
	"SA-final/book_microservice/proto/users"
	"SA-final/book_microservice/repository"
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

func (s *Server) AddBook(ctx context.Context, req *users.UserAddBookRequest) (*users.UserAddBookResponse, error) {
	bookName := req.GetBookName()
	book, err := s.app.bookRepo.SearchBookWithName(bookName)
	if err != nil {
		 return nil, err
	}
	res := &users.UserAddBookResponse{
		Book: &users.Book{
			Id:     int32(book.ID),
			Name:   book.Name,
			Author: book.Author,
		},
	}
	return res, err
}

func (s *Server) GetAllBooks(ctx context.Context, req *users.GetAllBooksRequest) (*users.GetAllBooksResponse, error) {
	books, err := s.app.bookRepo.GetAllBooks()
	if err != nil {
		return nil, err
	}
	protoBooks := []*users.Book{}
	for _, book := range books {
		protoBook := &users.Book{
			Id:     int32(book.ID),
			Name:   book.Name,
			Author: book.Author,
		}
		protoBooks = append(protoBooks, protoBook)
	}
	
	res := &users.GetAllBooksResponse{
		Books: protoBooks,
	}
	return res, err
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	bookRepo *repository.BookRepository
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := pgxpool.Connect(context.Background(), "user=postgres password=postgres host=db port=5432 dbname=postgres sslmode=disable pool_max_conns=10")
	if err != nil {
		errorLog.Fatal(err)
	}
	defer pool.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		bookRepo: &repository.BookRepository{Pool: pool},
	}

	server := &http.Server{
		Addr:     ":8000",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	go GRPCServe(app)

	infoLog.Print("Starting server on port 8000")
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func GRPCServe(app *application) {
	l, err := net.Listen("tcp", "booksapi:8001")
	if err != nil {
		log.Fatalf("Failed to listen:%v", err)
	}

	s := grpc.NewServer()
	users.RegisterUserServiceServer(s, &Server{app, users.UnimplementedUserServiceServer{}})

	log.Println("Server is running on port:8001")
	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve:%v", err)
	}
}
