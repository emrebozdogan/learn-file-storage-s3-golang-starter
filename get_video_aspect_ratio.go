package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Stream struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type FFprobeOutput struct {
	Streams []Stream `json:"streams"`
}

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)
	var cmdStdout bytes.Buffer

	cmd.Stdout = &cmdStdout

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	var ffprobeOutput FFprobeOutput
	err = json.NewDecoder(&cmdStdout).Decode(&ffprobeOutput)
	if err != nil {
		return "", err
	}

	if len(ffprobeOutput.Streams) == 0 {
		return "", fmt.Errorf("no streams found in video file")
	}

	stream := ffprobeOutput.Streams[0]
	width := stream.Width
	height := stream.Height

	if width <= 0 || height <= 0 {
		return "other", nil
	}

	ratio := float64(width) / float64(height)

	if ratio >= 1.7 && ratio <= 1.8 {
		return "16:9", nil
	} else if ratio >= 0.5 && ratio <= 0.6 {
		return "9:16", nil
	}

	return "other", nil
}
