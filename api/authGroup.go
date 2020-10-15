package api

import (
	"github.com/ridwanfathin/blog-server/api/handlers"

	// fixedwindow "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/fixedwindowcounter"
	// swl "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/slidingwindowlog"
	swc "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/slidingwindowcounter"

	"github.com/labstack/echo"
)

func AuthGroup(g *echo.Group) {
	// configAnon := fixedwindow.NewConfig("anonlimiter", 2, "minute")
	// g.Use(fixedwindow.AnonLimiter(configAnon))
	// configAnon := swl.NewConfig("anonlimiterswl", 5, "minute")
	// g.Use(swl.AnonLimiter(configAnon))
	configAnon := swc.NewConfig("anonlimiterswc", 5, "minute")
	g.Use(swc.AnonLimiter(configAnon))

	g.POST("/user", handlers.SignupUser)
	g.POST("/user/login", handlers.LoginUser)
}
