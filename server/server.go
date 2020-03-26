package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/config"
)

func Run() {
	fileWriter, err := os.Create("access.log")
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = fileWriter
	config.Load()
	fmt.Printf(" \n\t Listening [::]:%d \n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	g := *gin.Default()
	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "X-Request", "Location"},
		MaxAge:       1 * time.Minute,
	}))
}
