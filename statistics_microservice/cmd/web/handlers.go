package main

import (
	"SA-final/statistics_microservice/service"
	"encoding/json"
	"net/http"
)

func (app *application) statistic(w http.ResponseWriter, r *http.Request) {
	userList := app.getAllUsers(app.userConn)
	bookList := app.getAllBooks(app.bookConn)
	statistics := service.CalculateStatistic(userList, bookList)
	m, err := json.Marshal(statistics)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(m)
}
