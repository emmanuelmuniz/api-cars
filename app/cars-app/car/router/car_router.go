package router

import (
	controller "api-cars/app/app-controllers"

	"github.com/labstack/echo/v4"
)

func NewCarRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/cars", func(context echo.Context) error { return c.Car.GetCars(context) })
	e.GET("/cars/:id", func(context echo.Context) error { return c.Car.GetCar(context, context.Param("id")) })
	e.POST("/cars", func(context echo.Context) error { return c.Car.CreateCar(context) })
	e.DELETE("/cars/:id", func(context echo.Context) error { return c.Car.DeleteCar(context, context.Param("id")) })
	e.PUT("/cars", func(context echo.Context) error { return c.Car.UpdateCar(context) })
}
