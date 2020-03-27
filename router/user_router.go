package router

import (
	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/controllers"
)

func initRouterUser(e *gin.Engine) {
	groupE := e.Group("/api/v1/user")
	{
		groupE.POST("", controllers.CreateUser)
	}
}
