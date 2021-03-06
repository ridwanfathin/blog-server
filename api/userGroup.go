package api

import (
	"github.com/ridwanfathin/blog-server/api/handlers"

	fwc "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/fixedwindowcounter"

	"github.com/labstack/echo"
)

func UserGroup(g *echo.Group) {

	config := fwc.NewConfig("userlimiter", 5, "minute")
	g.Use(fwc.UserLimiter(config))

	g.GET("/users", handlers.GetAllUser)
	g.GET("/user/:id", handlers.GetUserByID)
	g.DELETE("/user/:id", handlers.DeleteUserByID)
	g.PUT("/user/:id", handlers.UpdateUser)
}
