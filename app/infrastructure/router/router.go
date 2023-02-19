package router

import (
	"api-cars/app/interface/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/cars", func(context echo.Context) error { return c.Car.GetCars(context) })
	e.GET("/cars/:id", func(context echo.Context) error { return c.Car.GetCar(context, context.Param("id")) })
	e.POST("/cars", func(context echo.Context) error { return c.Car.CreateCar(context) })
	e.DELETE("/cars/:id", func(context echo.Context) error { return c.Car.DeleteCar(context, context.Param("id")) })
	e.PUT("/cars", func(context echo.Context) error { return c.Car.UpdateCar(context) })
	return e
}
