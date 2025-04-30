package config

import (
	"log"
	"os"
)

// Config t.b.d. until API stable
type Config struct {
	Env          string
	CdnURL       string
	GCSBucket    string
	PostBasePath string
}

// New t.b.d. until API stable
func New() *Config {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
		log.Printf("warn: ENV env variable not set, defaulting to %s", env)
	}
	isDev := env == "dev"

	cdnURL := os.Getenv("CDN_URL")
	gcsBucket := os.Getenv("GCS_BUCKET")
	postBasePath := os.Getenv("POSTS_BASE_PATH")

	if cdnURL == "" && !isDev {
		log.Fatalf("error: CDN_URL env variable is required in higher environments")
	}
	if cdnURL == "" {
		cdnURL = "http://localhost:8080/public"
		log.Printf("warn: CDN_URL env variable not set, defaulting to %s", cdnURL)
	}

	if gcsBucket == "" && !isDev {
		log.Fatalf("error: GCS_BUCKET env variable is required in higher environments")
	}

	if postBasePath == "" {
		postBasePath = "./public/posts"
		log.Printf("warn: POSTS_BASE_PATH env variable not set, defaulting to %s", postBasePath)
	}

	return &Config{
		Env:          env,
		CdnURL:       cdnURL,
		GCSBucket:    gcsBucket,
		PostBasePath: postBasePath,
	}
}
