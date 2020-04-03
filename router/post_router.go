package router

import (
	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/controllers"
)

func initRouterPosts(g *gin.Engine) {
	groupP := g.Group("/api/v1/post")
	{
		groupP.GET("", controllers.GetPost)
		groupP.POST(":id", controllers.CreatePost)
		groupP.PUT(":id", controllers.UpdatePost)
	}
}
