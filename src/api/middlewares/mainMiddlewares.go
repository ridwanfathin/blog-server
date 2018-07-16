package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "../../../static",
	}))

	e.Use(serverHeader)
}

func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Shade/1.0")
		c.Response().Header().Set("customHeader", "thisHaveNoMeaning")

		return next(c)
	}
}
