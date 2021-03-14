package main

import (
	"github.com/bmizerany/pat"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := pat.New()

	router.Get("/api/v1/statistics", http.HandlerFunc(app.statistic))

	return router
}
