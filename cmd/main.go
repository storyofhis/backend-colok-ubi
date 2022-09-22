package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", func(ctx *gin.Context) {
		fmt.Println("Hello")
	})

	router.Run(":4000")
}
