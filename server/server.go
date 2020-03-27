package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/matkinhig/developer-community/config"
	"github.com/matkinhig/developer-community/router"
)

func Run(g *gin.Engine) {
	fileWriter, err := os.Create("access.log")
	if err != nil {
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = fileWriter
	config.Load()
	fmt.Printf(" \n\t Listening [::]:%d \n", config.PORT)
	router.New(g)
	listen(config.PORT, g)
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.PORT),
		Handler: g,
	}
	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func listen(port int, g *gin.Engine) {
	g.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{"Content-Type", "X-Request", "Location"},
		MaxAge:       12 * time.Hour,
	}))
}
