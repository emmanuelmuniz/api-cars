package router

import (
	controller "api-cars/app/app-controllers"

	"github.com/labstack/echo/v4"
)

func NewFeatureRouter(e *echo.Echo, c controller.AppController) {
	e.GET("/features", func(context echo.Context) error { return c.Feature.GetFeatures(context) })
	e.GET("/features/:id", func(context echo.Context) error { return c.Feature.GetFeature(context, context.Param("id")) })
	e.POST("/features", func(context echo.Context) error { return c.Feature.CreateFeature(context) })
	e.DELETE("/features/:id", func(context echo.Context) error { return c.Feature.DeleteFeature(context, context.Param("id")) })
	e.PUT("/features", func(context echo.Context) error { return c.Feature.UpdateFeature(context) })
}
