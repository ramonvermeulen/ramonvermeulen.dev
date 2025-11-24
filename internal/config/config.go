package config

import (
	"fmt"
	"log"
	"os"
)

// Config holds application runtime configuration.
type Config struct {
	Env          string
	CdnURL       string
	GCSBucket    string
	PostBasePath string
	BaseURL      string
}

// New constructs a Config from environment variables, applying defaults.
// Returns an error instead of exiting on validation failures.
func New() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
		log.Printf("warn: ENV env variable not set, defaulting to %s", env)
	}
	isDev := env == "dev"

	cdnURL := os.Getenv("CDN_URL")
	gcsBucket := os.Getenv("GCS_BUCKET")
	postBasePath := os.Getenv("POSTS_BASE_PATH")
	baseURL := os.Getenv("BASE_URL")

	if cdnURL == "" && !isDev {
		return nil, fmt.Errorf("CDN_URL env variable is required when ENV=%s", env)
	}
	if cdnURL == "" { // dev default
		cdnURL = "http://localhost:8080/public"
		log.Printf("warn: CDN_URL env variable not set, defaulting to %s", cdnURL)
	}
	if gcsBucket == "" && !isDev {
		return nil, fmt.Errorf("GCS_BUCKET env variable is required when ENV=%s", env)
	}
	if postBasePath == "" {
		postBasePath = "./public/posts/"
		log.Printf("warn: POSTS_BASE_PATH env variable not set, defaulting to %s", postBasePath)
	}
	if baseURL == "" {
		baseURL = "http://localhost:8080"
		log.Printf("warn: BASE_URL env variable not set, defaulting to %s", baseURL)
	}

	return &Config{
		Env:          env,
		CdnURL:       cdnURL,
		GCSBucket:    gcsBucket,
		PostBasePath: postBasePath,
		BaseURL:      baseURL,
	}, nil
}
