package router

import (
	controller "api-cars/app/app-controllers"

	bsr "api-cars/app/cars-app/body-style/router"
	cmr "api-cars/app/cars-app/car-model/router"
	cr "api-cars/app/cars-app/car/router"
	fr "api-cars/app/cars-app/feature/router"
	mr "api-cars/app/cars-app/make/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bsr.NewBodyStyleRouter(e, c)
	mr.NewMakeRouter(e, c)
	cmr.NewCarModelRouter(e, c)
	cr.NewCarRouter(e, c)
	fr.NewFeatureRouter(e, c)

	return e
}
