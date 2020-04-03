package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPost(g *gin.Context) {
	fmt.Println("Get Post")
}

func CreatePost(g *gin.Context) {
	fmt.Println("Create post")
}

func UpdatePost(g *gin.Context) {
	fmt.Println("Update Post")
}
