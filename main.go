package main

import (
	"video_convertor/internal/api"
	"video_convertor/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	var (
		videoService = service.NewVideoService()
		videoHandler = api.NewVideoHandler(videoService)
	)

	// define new router
	router := gin.Default()

	router.POST("/upload", videoHandler.HandleUpload)
	router.GET("/download/:fileID", videoHandler.HandleDownload)

	// start server
	router.Run("localhost:3000")
}
