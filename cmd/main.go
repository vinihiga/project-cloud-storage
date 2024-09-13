package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinihiga/project-cloud-storage/internal/app/handlers"
)

func main() {
	var server *gin.Engine = gin.Default()
	server.MaxMultipartMemory = 8 << 20 // 8 MiB

	server.GET("/", handlers.ListHandler)
	server.GET("/list", handlers.ListHandler)
	server.POST("/upload", handlers.UploadHandler)
	server.GET("/download/:filename", handlers.DownloadHandler)

	server.RunTLS(":8080", "server.pem", "server.key")
}
