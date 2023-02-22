package router

import (
	"api-cars/app/interface/controller"

	"github.com/labstack/echo/v4"
)

func NewBodyStyleRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/body-styles", func(context echo.Context) error { return c.BodyStyle.GetBodyStyles(context) })
	e.GET("/body-styles/:id", func(context echo.Context) error { return c.BodyStyle.GetBodyStyle(context, context.Param("id")) })
	e.POST("/body-styles", func(context echo.Context) error { return c.BodyStyle.CreateBodyStyle(context) })
	e.DELETE("/body-styles/:id", func(context echo.Context) error { return c.BodyStyle.DeleteBodyStyle(context, context.Param("id")) })
	e.PUT("/body-styles", func(context echo.Context) error { return c.BodyStyle.UpdateBodyStyle(context) })
}
