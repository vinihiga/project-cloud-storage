package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vinihiga/project-cloud-storage/internal/app/utils"
)

func UploadHandler(context *gin.Context) {
	file, _ := context.FormFile("file")
	context.SaveUploadedFile(file, utils.PublicPath+file.Filename)
	context.Done()
}
