package router

import (
	"api-cars/app/interface/controller"

	"github.com/labstack/echo/v4"
)

func NewCarModelRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/car-models", func(context echo.Context) error { return c.CarModel.GetCarModels(context) })
	e.GET("/car-models/:id", func(context echo.Context) error { return c.CarModel.GetCarModel(context, context.Param("id")) })
	e.POST("/car-models", func(context echo.Context) error { return c.CarModel.CreateCarModel(context) })
	e.DELETE("/car-models/:id", func(context echo.Context) error { return c.CarModel.DeleteCarModel(context, context.Param("id")) })
	e.PUT("/car-models", func(context echo.Context) error { return c.CarModel.UpdateCarModel(context) })
}
