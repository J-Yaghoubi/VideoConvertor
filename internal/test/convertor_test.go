package test

import (
	"testing"
	"video_convertor/internal/config"
	"video_convertor/internal/helper"
)

func BenchmarkEncodeVideo(b *testing.B) {
	// take benchmark from converting process
	// note: fill the inputs with valid input file and then execute test with:
	// go test -bench=. -run=^$

	var des = "output" + config.DESTINATION_FORMAT

	for i := 0; i < b.N; i++ {
		helper.EncodeVideo("input.mp4", des, config.DESTINATION_CODEC, "hd")
	}
}
