package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/vinihiga/project-cloud-storage/internal/app/utils"
)

type listResponse struct {
	Files []string `json:"files"`
}

func ListHandler(context *gin.Context) {
	files, err := os.ReadDir(utils.PublicPath)

	if err != nil {
		log.Fatal(err)
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var response = listResponse{Files: make([]string, 0)}

	for _, file := range files {
		if !file.IsDir() {
			response.Files = append(response.Files, file.Name())
		}
	}

	context.JSON(http.StatusOK, response)
}
