package markdown

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

// GCSReader reads files from Google Cloud Storage
type GCSReader struct {
	BucketName string
	Client     *storage.Client
}

// Open returns a ReadCloser for the specified target file in GCS
func (gr *GCSReader) Open(target string) (io.ReadCloser, error) {
	ctx := context.Background()
	reader, err := gr.Client.Bucket(gr.BucketName).Object(target).NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, ErrFileNotFound
		}
		return nil, ErrReadFailed
	}
	return reader, nil
}

// List returns a list of file names in GCS with the specified prefix
func (gr *GCSReader) List(prefix string) ([]string, error) {
	ctx := context.Background()
	it := gr.Client.Bucket(gr.BucketName).Objects(ctx, &storage.Query{Prefix: prefix})
	var fileNames []string
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error listing files: %w", err)
		}
		fileNames = append(fileNames, attrs.Name)
	}
	return fileNames, nil
}
