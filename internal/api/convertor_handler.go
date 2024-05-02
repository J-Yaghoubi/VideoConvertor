package api

import (
	"fmt"
	"net/http"

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
		respondWithError(c, http.StatusBadRequest, fmt.Sprintf("error retrieving the video: %v", err))
		return
	}

	result, err := vh.svc.Process(c, file, "sd")
	if err != nil {
		respondWithError(c, http.StatusBadRequest, fmt.Sprintf("%v", err))
		return
	}

	respondWithJSON(c, 201, result)
	return
}

func (vh VideoHandler) HandleDownload(c *gin.Context) {
	// download converted file if exists

	fileID := c.Param("fileID")

	if !helper.IsValidPath(fileID) {
		respondWithError(c, http.StatusBadRequest, fmt.Sprint("Invalid File ID format"))
		return
	}

	if !helper.IsFileExists("./uploads/" + fileID) {
		respondWithError(c, http.StatusNotFound, fmt.Sprint("Not Found!"))
		return
	}

	respondWithFile(c, http.StatusAccepted, fileID)
	return
}
