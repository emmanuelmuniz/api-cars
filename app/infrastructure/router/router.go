package router

import (
	"api-cars/app/interface/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	NewBodyStyleRouter(e, c)
	NewMakeRouter(e, c)
	NewCarModelRouter(e, c)
	NewCarRouter(e, c)

	return e
}
