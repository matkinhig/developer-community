package router

import (
	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/controllers"
)

func UserRouter(g *gin.Engine) {
	e := g.Group("/api/v1")
	e.POST("/user", controllers.CreateUser())
}
