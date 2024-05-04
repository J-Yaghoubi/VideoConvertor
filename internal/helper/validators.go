package helper

import (
	"mime/multipart"
	"os"
	"regexp"
)

func IsFileExists(path string) bool {
	// check if given path is exists or not

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func IsValidPath(path string) bool {
	// check if given path has not invalid characters

	var validFileNamePattern = regexp.MustCompile(`^[\w\-\._\d]+$`)
	return validFileNamePattern.MatchString(path)
}

func IsValidateFileType(file *multipart.FileHeader) bool {
	// validate file mime type to ensure it is video

	if file.Header.Get("Content-Type") != "video/mp4" {
		return false
	}

	return true
}

func IsValidateSize(file *multipart.FileHeader) bool {
	// validate file size

	var maxFileSize int64 = 100_000_000

	if file.Size > maxFileSize {
		return false
	}

	return true
}

func IsValidateQuality(quality string) bool {
	// Validate the quality value

	validQuality := false

	switch quality {
	case "hd", "fhd", "sd":
		validQuality = true
	default:
		validQuality = false
	}

	if !validQuality {
		return false
	}

	return true
}
