package router

import (
	controller "api-cars/app/app-controllers"

	"github.com/labstack/echo/v4"
)

func NewCarImageRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/car-images", func(context echo.Context) error { return c.CarImage.GetCarImages(context) })
	e.GET("/car-images/:id", func(context echo.Context) error { return c.CarImage.GetCarImage(context, context.Param("id")) })
	e.POST("/car-images", func(context echo.Context) error { return c.CarImage.CreateCarImage(context) })
	e.DELETE("/car-images/:id", func(context echo.Context) error { return c.CarImage.DeleteCarImage(context, context.Param("id")) })
	e.PUT("/car-images", func(context echo.Context) error { return c.CarImage.UpdateCarImage(context) })
}
