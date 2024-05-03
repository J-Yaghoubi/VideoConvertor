package api

import (
	"context"
	"net/http"
	"time"

	"video_convertor/views"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
)

const appTimeout = time.Second * 10

func render(c *gin.Context, status int, template templ.Component) error {
	c.Status(status)
	return template.Render(c.Request.Context(), c.Writer)
}

func HandleIndexView(c *gin.Context) {
	// upload a video file to start converting process

	_, cancel := context.WithTimeout(context.Background(), appTimeout)
	defer cancel()

	render(c, http.StatusOK, views.Index())
}
