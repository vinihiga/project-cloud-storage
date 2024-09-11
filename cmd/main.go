package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var server *gin.Engine = gin.Default()

	server.GET("/download", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.RunTLS(":8080", "server.pem", "server.key")
}
