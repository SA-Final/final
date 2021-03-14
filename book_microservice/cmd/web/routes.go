package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := pat.New()

	router.Post("/api/v1/books", http.HandlerFunc(app.createBook))
	router.Get("/api/v1/books", http.HandlerFunc(app.getAllBooks))
	router.Get("/api/v1/books/:id", http.HandlerFunc(app.getBookWithId))
	router.Post("/api/v1/books/search", http.HandlerFunc(app.searchBook))

	return router
}
