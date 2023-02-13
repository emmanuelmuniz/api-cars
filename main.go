package main

import (
	"net/http"

	"github.com/emmanuelmuniz/api-cars/db"
	"github.com/emmanuelmuniz/api-cars/models"
	"github.com/emmanuelmuniz/api-cars/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.Car{})

	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
