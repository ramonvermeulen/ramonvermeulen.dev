package config

import (
	"log"
	"os"
)

// Config t.b.d. until API stable
type Config struct {
	Environment  string
	AssetURL     string
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
	isProd := environment == "production"

	assetURL := os.Getenv("ASSET_URL")
	if assetURL == "" && isProd {
		log.Fatalf("error: ASSET_URL environment variable is required in production")
	} else {
		assetURL = "./public"
	}

	gcsBucket := os.Getenv("GCS_BUCKET")
	if gcsBucket == "" && isProd {
		log.Fatalf("error: GCS_BUCKET environment variable is required in production")
	}

	postBasePath := os.Getenv("POSTS_BASE_PATH")
	if postBasePath == "" {
		postBasePath = "./public/posts"
		log.Printf("warn: POSTS_BASE_PATH environment variable not set, defaulting to %s", postBasePath)
	}

	return &Config{
		Environment:  environment,
		AssetURL:     assetURL,
		GCSBucket:    gcsBucket,
		PostBasePath: postBasePath,
	}
}
