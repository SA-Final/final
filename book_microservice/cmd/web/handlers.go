package main

import (
	"SA-final/book_microservice/models"
	"encoding/json"
	"net/http"
	"strconv"
)

type bookName struct {
	Name string `json:"name"`
}

func (app *application) createBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	book := models.Book{}
	err := decoder.Decode(&book)
	if err != nil {
		app.serverError(w, err)
		return
	}

	createdBook, err := app.bookRepo.CreateBook(&book)
	if err != nil {
		app.serverError(w, err)
		return
	}

	m, err := json.Marshal(createdBook)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) getAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := app.bookRepo.GetAllBooks()
	if err != nil {
		app.serverError(w, err)
		return
	}
	m, err := json.Marshal(books)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) getBookWithId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	book, err := app.bookRepo.GetBookWithId(id)
	if err != nil {
		app.notFound(w)
		return
	}

	m, err := json.Marshal(book)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(m)
}

func (app *application) searchBook(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	name := bookName{}

	err := decoder.Decode(&name)
	if err != nil {
		app.serverError(w, err)
		return
	}
	books, err := app.bookRepo.SearchBook(name.Name)
	m, err := json.Marshal(books)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(m)
}
