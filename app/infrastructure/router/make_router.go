package router

import (
	"api-cars/app/interface/controller"

	"github.com/labstack/echo/v4"
)

func NewMakeRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/makes", func(context echo.Context) error { return c.Make.GetMakes(context) })
	e.GET("/makes/:id", func(context echo.Context) error { return c.Make.GetMake(context, context.Param("id")) })
	e.POST("/makes", func(context echo.Context) error { return c.Make.CreateMake(context) })
	e.DELETE("/makes/:id", func(context echo.Context) error { return c.Make.DeleteMake(context, context.Param("id")) })
	e.PUT("/makes", func(context echo.Context) error { return c.Make.UpdateMake(context) })
}
