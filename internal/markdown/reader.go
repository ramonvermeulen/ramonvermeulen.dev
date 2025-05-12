package markdown

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
)

// FileReader t.b.d. until API stable
type FileReader interface {
	Open(target string) (io.ReadCloser, error)
	List(prefix string) ([]string, error)
}

// NewFileReader t.b.d. until API stable
func NewFileReader(cfg *config.Config) (FileReader, error) {
	if cfg.Env != "dev" {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		return &GCSReader{
			BucketName: cfg.GCSBucket,
			Client:     client,
		}, nil
	}
	return &LocalReader{}, nil
}
