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

	e.GET("/makes", func(context echo.Context) error { return c.Make.GetMakes(context) })
	e.GET("/makes/:id", func(context echo.Context) error { return c.Make.GetMake(context, context.Param("id")) })
	e.POST("/makes", func(context echo.Context) error { return c.Make.CreateMake(context) })
	e.DELETE("/makes/:id", func(context echo.Context) error { return c.Make.DeleteMake(context, context.Param("id")) })
	e.PUT("/makes", func(context echo.Context) error { return c.Make.UpdateMake(context) })

	e.GET("/car-models", func(context echo.Context) error { return c.CarModel.GetCarModels(context) })
	e.GET("/car-models/:id", func(context echo.Context) error { return c.CarModel.GetCarModel(context, context.Param("id")) })
	e.POST("/car-models", func(context echo.Context) error { return c.CarModel.CreateCarModel(context) })
	e.DELETE("/car-models/:id", func(context echo.Context) error { return c.CarModel.DeleteCarModel(context, context.Param("id")) })
	e.PUT("/car-models", func(context echo.Context) error { return c.CarModel.UpdateCarModel(context) })
	return e
}
