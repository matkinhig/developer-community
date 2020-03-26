package responses

import "github.com/gin-gonic/gin"

func FAIL(g *gin.Context, code int, err error) {
	g.AbortWithStatusJSON(code, gin.H{
		"Errors": err.Error(),
	})
}

func DONE(g *gin.Context, code int, data interface{}) {
	g.Writer.WriteHeader(code)
	g.JSON(code, data)
}
