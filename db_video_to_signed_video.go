package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"
)

func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {
	var splittedVideoURL []string
	videoURL := video.VideoURL

	if videoURL == nil {
		return video, nil
	}

	splittedVideoURL = strings.Split(*videoURL, ",")

	if len(splittedVideoURL) != 2 {
		return video, fmt.Errorf("invalid video URL format")
	}

	bucket := splittedVideoURL[0]
	key := splittedVideoURL[1]

	presignedURL, err := generatePresignedURL(cfg.s3Client, bucket, key, 5*time.Minute)
	if err != nil {
		return video, fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	video.VideoURL = &presignedURL

	return video, nil
}
