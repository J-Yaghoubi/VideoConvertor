package service

import (
	"errors"
	"mime/multipart"
	"video_convertor/internal/helper"
	"video_convertor/internal/port"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type videoService struct{}

func NewVideoService() port.VideoService {
	return &videoService{}
}

func (vs *videoService) Process(c *gin.Context, file *multipart.FileHeader, resolution string) (*string, error) {
	// video converting process

	if !helper.IsValidateType(file) {
		return nil, errors.New("Invalid file type")
	}

	if !helper.IsValidateSize(file) {
		return nil, errors.New("File size exceeds the limit")
	}

	directory := "./uploads/"
	fileExt := helper.GetFileExtension(file.Filename)
	tempName := uuid.New().String() + fileExt
	tempPath := directory + tempName
	cvName := uuid.New().String() + fileExt
	cvPath := directory + cvName

	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		return nil, errors.New("Failed to save file")
	}

	go helper.EncodeVideo(tempPath, cvPath, "libx264", resolution)

	return &cvName, nil
}
