package markdown

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/storage"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
)

var (
	// ErrFileNotFound indicates that the requested file does not exist
	ErrFileNotFound = errors.New("file not found")
	// ErrReadFailed indicates that reading the file failed
	ErrReadFailed = errors.New("failed to read file")
)

// FileReader interface for reading and listing files
type FileReader interface {
	Open(target string) (io.ReadCloser, error)
	List(prefix string) ([]string, error)
}

// NewFileReader returns a new FileReader based on the environment
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
