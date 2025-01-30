package markdown

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/storage"
)

// GCSReader t.b.d. until API stable
type GCSReader struct {
	BucketName string
	Client     *storage.Client
}

// Read t.b.d. until API stable
func (gr *GCSReader) Read(target string) ([]byte, error) {
	ctx := context.Background()
	rc, err := gr.Client.Bucket(gr.BucketName).Object(target).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, ErrFileNotFound
		}
		return nil, ErrReadFailed
	}
	defer rc.Close()
	content, err := io.ReadAll(rc)
	if err != nil {
		return nil, ErrReadFailed
	}
	return content, nil
}
