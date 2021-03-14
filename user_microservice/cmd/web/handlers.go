package main

import (
	"SA-final/user_microservice/models"
	"encoding/json"
	"fmt"
	"net/http"
)

type AddBook struct {
	BookName string `json:"bookName"`
	Username string `json:"username"`
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	err := decoder.Decode(&user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	newUser, err := app.userRepo.RegisterUser(&user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	m, err := json.Marshal(newUser)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) addBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	addBook := &AddBook{}
	err := decoder.Decode(&addBook)
	if err != nil {
		app.serverError(w, err)
		return
	}

	user, err := app.userRepo.GetUserByUsername(addBook.Username)
	fmt.Print(user)
	if err != nil {
		app.serverError(w, err)
	}
	book := app.getBook(app.conn, addBook.BookName)
	user, err = app.userRepo.AddBook(*user, *book)

	if err != nil {
		app.serverError(w, err)
		return
	}

	m, err := json.Marshal(user)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.userRepo.GetAllUsers()
	if err != nil {
		app.serverError(w, err)
		return
	}

	m, err := json.Marshal(users)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) getUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get(":username")
	user, err := app.userRepo.GetUserByUsername(username)
	if err != nil {
		app.notFound(w)
		return
	}
	m, err := json.Marshal(user)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}
