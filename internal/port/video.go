package port

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type VideoService interface {
	Process(ctx *gin.Context, file *multipart.FileHeader, quality string) (*string, error)
}
