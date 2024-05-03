package service

import (
	"errors"
	"mime/multipart"
	"video_convertor/internal/config"
	"video_convertor/internal/helper"
	"video_convertor/internal/port"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type videoService struct{}

func NewVideoService() port.VideoService {
	return &videoService{}
}

func (vs *videoService) Process(c *gin.Context, file *multipart.FileHeader, quality string) (*string, error) {
	// video converting process

	if !helper.IsValidateFileType(file) {
		return nil, errors.New("Invalid file type")
	}

	if !helper.IsValidateSize(file) {
		return nil, errors.New("File size exceeds the limit")
	}

	if !helper.IsValidateQuality(quality) {
		return nil, errors.New("Invalid quality")
	}

	fileExt := helper.GetFileExtension(file.Filename)
	tempName := uuid.New().String() + fileExt
	tempPath := config.UPLOAD_DIRECTORY + tempName
	cvName := uuid.New().String() + config.DESTINATION_FORMAT
	cvPath := config.UPLOAD_DIRECTORY + cvName

	if err := c.SaveUploadedFile(file, tempPath); err != nil {
		return nil, errors.New("Failed to save file")
	}

	go helper.EncodeVideo(tempPath, cvPath, config.DESTINATION_CODEC, quality)

	return &cvName, nil
}
