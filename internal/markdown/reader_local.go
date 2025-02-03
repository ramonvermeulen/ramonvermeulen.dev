package markdown

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LocalReader t.b.d. until API stable
type LocalReader struct{}

// Open t.b.d. until API stable
func (lr *LocalReader) Open(target string) (io.ReadCloser, error) {
	file, err := os.Open(target)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, ErrReadFailed
	}
	return file, nil
}

// List t.b.d. until API stable
func (lr *LocalReader) List(prefix string) ([]string, error) {
	fileNames, err := filepath.Glob(fmt.Sprintf("%s/*.md", prefix))
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}
	return fileNames, nil
}
