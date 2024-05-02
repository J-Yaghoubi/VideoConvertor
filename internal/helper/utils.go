package helper

import (
	"os"
	"os/exec"
	"strings"
)

func GetFileExtension(filename string) string {
	// extract and return file extension from filename

	split := strings.Split(filename, ".")
	if len(split) > 0 && split[len(split)-1] != "" {
		return "." + split[len(split)-1]
	}
	return ""
}

func EncodeVideo(input, output, codec string, resolution string) error {
	// convert given file to requested format and delete original one

	res := map[string]string{
		"fhd": "1920x1080",
		"hd":  "1280x720",
		"sd":  "640x480",
	}

	var dimension string

	if val, ok := res[resolution]; ok {
		dimension = val
	} else {
		dimension = res["sd"]
	}

	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", codec, "-s", dimension, output)
	err := cmd.Run()
	if err != nil {
		return err
	}

	os.Remove(input)
	return nil
}
