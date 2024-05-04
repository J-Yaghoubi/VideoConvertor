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
		viewHandler  = api.NewViewHandler()
	)

	// define new router
	router := gin.Default()

	// api routes
	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/upload", videoHandler.HandleUpload)
		apiGroup.GET("/download/:fileID", videoHandler.HandleDownload)
	}

	// web ui routes
	webGroup := router.Group("/")
	{
		webGroup.GET("/", viewHandler.HandleIndexView)
	}

	// start server
	router.Run(config.SERVER_ADDRESS)
}
