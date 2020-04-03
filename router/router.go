package router

import (
	"github.com/gin-gonic/gin"
)

func New(e *gin.Engine) {
	e.GET("/ping", pingHealth)
	initRouterUser(e)
	initRouterPosts(e)

}
