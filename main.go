package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/server"
)

func main() {
	fmt.Println("start golang...")
	g := gin.Default()
	server.Run(g)
}
