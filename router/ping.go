package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/config"
)

var counter int = 0

func pingHealth(c *gin.Context) {
	counter++
	// Race condition => Hai CPU cung access vo 1 cai bien
	// Atomic => co 100 request vao nhung expected: counter = 100, actual: 80
	c.Writer.Header().Set("X-Counter", strconv.Itoa(counter))
	message := "Ping from "
	config.Load()
	port := config.PORT
	if port == 3456 {
		message = message + "1nd server"
	} else {
		message = message + "2nd server"
	}
	c.String(200, message)
}
