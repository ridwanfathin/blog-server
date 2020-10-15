package api

import (
	"github.com/ridwanfathin/blog-server/api/handlers"
	swl "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/slidingwindowlog"

	"github.com/labstack/echo"
)

func SwlGroup(g *echo.Group) {

	config := swl.NewConfig("userlimiterswl", 5, "minute")
	g.Use(swl.UserLimiter(config))

	g.POST("/publish", handlers.PublishPost)
	g.GET("/posts", handlers.GetAllPost)
	g.GET("/post/:id", handlers.GetPostByID)
	g.GET("/author/:id/posts", handlers.GetPostByAuthorID)
	g.DELETE("/post/:id", handlers.DeletePostByID)
	g.PUT("/post/:id", handlers.UpdatePost)
}
