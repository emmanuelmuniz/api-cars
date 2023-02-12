package main

import (
	"net/http"

	"github.com/emmanuelmuniz/api-cars/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", router)
}
