package api

import (
	"context"
	"net/http"
	"time"

	"video_convertor/views"

	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

type viewHandler struct{}

func NewViewHandler() *viewHandler {
	return &viewHandler{}
}

func (vh *viewHandler) HandleIndexView(c *gin.Context) {
	// handle index page

	_, cancel := context.WithTimeout(context.Background(), appTimeout)
	defer cancel()

	respondWithTemplate(c, http.StatusOK, views.Index())
}
