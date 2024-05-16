package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{"hello": "world"})
	})
	router.Run("localhost:3000")
}
