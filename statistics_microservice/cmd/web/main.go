package main

import (
	"SA-final/statistics_microservice/models"
	"SA-final/statistics_microservice/proto/users"
	"context"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	bookConn users.UserServiceClient
	userConn users.UserServiceClient
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		bookConn: GRPC("booksapi:8001"),
		userConn: GRPC("usersapi:8003"),
	}

	server := &http.Server{
		Addr:     ":8004",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Print("Starting server on port 8004")
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}

func GRPC(url string) users.UserServiceClient {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return users.NewUserServiceClient(conn)
}

func (app *application) getAllBooks(c users.UserServiceClient) []*models.Book {
	ctx := context.Background()
	request := &users.GetAllBooksRequest{}
	response, err := c.GetAllBooks(ctx, request)
	if err != nil {
		log.Fatalf("error while calling GetAllBooks RPC %v", err)
	}
	books := []*models.Book{}
	for _, protoBook := range response.Books {
		book := &models.Book{
			ID:     int(protoBook.Id),
			Name:   protoBook.Name,
			Author: protoBook.Author,
		}
		books = append(books, book)
	}
	return books
}

func (app *application) getAllUsers(c users.UserServiceClient) []*models.User{
	ctx := context.Background()
	request := &users.GetAllUsersRequest{}
	response, err := c.GetAllUsers(ctx, request)
	if err != nil {
		log.Fatalf("error while calling GetAllUsers RPC %v", err)
	}
	userList := []*models.User{}
	for _, protoUser := range response.Users {
		bookIds := []int{}
		for _, id := range protoUser.BookIds {
			bookIds = append(bookIds, int(id))
		}
		user := &models.User{
			ID:       int(protoUser.Id),
			Email:    protoUser.Email,
			Username: protoUser.Username,
			Password: protoUser.Password,
			BookIds:  bookIds,
		}
		userList = append(userList, user)
	}
	return userList
}
