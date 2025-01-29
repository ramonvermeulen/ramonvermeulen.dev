package markdown

import (
	"context"
	"errors"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type GCSReader struct {
	BucketName string
	Client     *storage.Client
	BasePath   string
}

func (gr *GCSReader) Read(target string) ([]byte, error) {
	ctx := context.Background()
	objectPath := fmt.Sprintf("%s/%s.md", gr.BasePath, target)
	rc, err := gr.Client.Bucket(gr.BucketName).Object(objectPath).NewReader(ctx)
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
