package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"video_convertor/internal/config"

	"github.com/gin-gonic/gin"
)

func respondWithFile(c *gin.Context, code int, fileID string) {
	// send file response

	byteFile, err := os.ReadFile(config.UPLOAD_DIRECTORY + fileID)
	if err != nil {
		fmt.Println(err)
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileID))
	c.Data(code, "application/octet-stream", byteFile)

}

func respondWithText(c *gin.Context, code int, msg string) {
	// send text response

	c.String(code, msg)
}

func respondWithJSON(c *gin.Context, code int, payload interface{}) {
	// send json response

	_, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response %v", payload)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(code, payload)
}

func respondWithError(c *gin.Context, code int, msg string) {
	// send error response

	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(c, code, errResponse{
		Error: msg,
	})
}
