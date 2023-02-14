package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/cars", GetAllCars).Methods("GET")
	router.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	router.HandleFunc("/cars", CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	return router
}
