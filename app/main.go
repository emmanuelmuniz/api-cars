package main

import (
	"api-cars/app/config"
	"api-cars/app/infrastructure/datastore"
	"api-cars/app/infrastructure/router"
	"api-cars/app/registry"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
