package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vinihiga/project-cloud-storage/internal/app/utils"
)

func DownloadHandler(context *gin.Context) {
	fileName := context.Param("filename")
	isSlashAvailable := strings.Contains(fileName, "/")
	isReverseSlashAvailable := strings.Contains(fileName, "\\")

	if isSlashAvailable || isReverseSlashAvailable {
		context.AbortWithStatus(http.StatusNotFound)
		return
	}

	context.File(utils.PublicPath + fileName)
}
