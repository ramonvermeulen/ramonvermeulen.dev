package config

import (
	"log"
	"os"
)

// Config t.b.d. until API stable
type Config struct {
	Environment  string
	CdnURL       string
	GCSBucket    string
	PostBasePath string
}

// New t.b.d. until API stable
func New() *Config {
	environment := os.Getenv("ENV")
	if environment == "" {
		environment = "development"
		log.Printf("warn: ENV environment variable not set, defaulting to %s", environment)
	}
	isDev := environment == "development"

	cdnURL := os.Getenv("CDN_URL")
	if cdnURL == "" && !isDev {
		log.Fatalf("error: CDN_URL environment variable is required in higher environments")
	} else {
		cdnURL = "./public"
		log.Printf("warn: CDN_URL environment variable not set, defaulting to %s", cdnURL)
	}

	gcsBucket := os.Getenv("GCS_BUCKET")
	if gcsBucket == "" && !isDev {
		log.Fatalf("error: GCS_BUCKET environment variable is required in higher environments")
	}

	postBasePath := os.Getenv("POSTS_BASE_PATH")
	if postBasePath == "" {
		postBasePath = "./public/posts"
		log.Printf("warn: POSTS_BASE_PATH environment variable not set, defaulting to %s", postBasePath)
	}

	return &Config{
		Environment:  environment,
		CdnURL:       cdnURL,
		GCSBucket:    gcsBucket,
		PostBasePath: postBasePath,
	}
}
