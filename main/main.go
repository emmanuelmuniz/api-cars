package main

import (
	"net/http"

	"github.com/emmanuelmuniz/api-cars/db"
	controller "github.com/emmanuelmuniz/api-cars/main/cars/controllers"
)

func main() {

	handler := controller.New()

	server := &http.Server{
		Addr:    "0.0.0.0:3000",
		Handler: handler,
	}

	db.DBConnection()

	server.ListenAndServe()
}
