package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := pat.New()

	router.Post("/api/v1/users", http.HandlerFunc(app.registerUser))
	router.Post("/api/v1/users/addBook", http.HandlerFunc(app.addBook))
	router.Get("/api/v1/users", http.HandlerFunc(app.getAllUsers))
	router.Get("/api/v1/users/:username", http.HandlerFunc(app.getUserByUsername))

	return router
}
