package api

import (
	"github.com/ridwanfathin/blog-server/api/handlers"
	swc "github.com/ridwanfathin/blog-server/api/middlewares/ratelimiter/slidingwindowcounter"

	"github.com/labstack/echo"
)

func SwcGroup(g *echo.Group) {

	config := swc.NewConfig("userlimiterswc", 5, "minute")
	g.Use(swc.UserLimiter(config))

	g.POST("/publish", handlers.PublishPost)
	g.GET("/posts", handlers.GetAllPost)
	g.GET("/post/:id", handlers.GetPostByID)
	g.GET("/author/:id/posts", handlers.GetPostByAuthorID)
	g.DELETE("/post/:id", handlers.DeletePostByID)
	g.PUT("/post/:id", handlers.UpdatePost)
}
