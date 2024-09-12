package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var server *gin.Engine = gin.Default()
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	server.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.SaveUploadedFile(file, "../tests/"+file.Filename)
		context.JSON(http.StatusOK, gin.H{})
	})

	server.GET("/download/:filename", func(context *gin.Context) {
		fileName := context.Param("filename")
		context.File("../tests/" + fileName)
	})

	server.RunTLS("localhost:8080", "server.pem", "server.key")
}
