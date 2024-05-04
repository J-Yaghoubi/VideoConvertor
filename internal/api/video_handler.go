package api

import (
	"fmt"
	"net/http"

	"video_convertor/internal/config"
	"video_convertor/internal/helper"
	"video_convertor/internal/port"

	"github.com/gin-gonic/gin"
)

type VideoHandler struct {
	svc port.VideoService
}

func NewVideoHandler(svc port.VideoService) *VideoHandler {
	return &VideoHandler{
		svc,
	}
}

func (vh VideoHandler) HandleUpload(c *gin.Context) {
	// upload a video file to start converting process

	file, err := c.FormFile("video")
	if err != nil {
		respondWithError(c, http.StatusBadRequest, "missing video parameter")
		return
	}

	quality, ok := c.GetPostForm("quality")
	if !ok {
		respondWithError(c, http.StatusBadRequest, "missing quality parameter")
		return
	}

	result, err := vh.svc.Process(c, file, quality)
	if err != nil {
		respondWithError(c, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	msg := fmt.Sprintf("http://%s/api/download/%s", config.SERVER_ADDRESS, *result)
	respondWithText(c, 201, msg)
	return
}

func (vh VideoHandler) HandleDownload(c *gin.Context) {
	// download converted file if exists

	fileID := c.Param("fileID")

	if !helper.IsValidPath(fileID) {
		respondWithError(c, http.StatusBadRequest, fmt.Sprint("Invalid File ID format"))
		return
	}

	if !helper.IsFileExists(config.UPLOAD_DIRECTORY + fileID) {
		respondWithError(c, http.StatusNotFound, fmt.Sprint("Not Found!"))
		return
	}

	respondWithFile(c, http.StatusAccepted, fileID)
	return
}
