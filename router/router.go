package router

import (
	"github.com/gin-gonic/gin"
)

func New(e *gin.Engine) {
	e.GET("/ping", pingHealth)
	// e := g.Group("/api/v1")
	// {
	// 	e.POST("/user", controllers.CreateUser)
	// }
	initRouterUser(e)
}
