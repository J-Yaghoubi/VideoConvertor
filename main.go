package main

import (
	"video_convertor/internal/api"
	"video_convertor/internal/config"
	"video_convertor/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {

	var (
		videoService = service.NewVideoService()
		videoHandler = api.NewVideoHandler(videoService)
		viewHandler  = api.HandleIndexView
	)

	// define new router
	router := gin.Default()

	router.GET("/", viewHandler)
	router.POST("/upload", videoHandler.HandleUpload)
	router.GET("/download/:fileID", videoHandler.HandleDownload)

	// start server
	router.Run(config.SERVER_ADDRESS)
}
