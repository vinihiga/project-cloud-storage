package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var server *gin.Engine = gin.Default()
	server.MaxMultipartMemory = 8 << 20

	server.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.SaveUploadedFile(file, "../tests/"+file.Filename)
		context.JSON(http.StatusOK, gin.H{})
	})

	server.GET("/download", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{})
	})

	server.RunTLS(":8080", "server.pem", "server.key")
}
