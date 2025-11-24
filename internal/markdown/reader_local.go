package markdown

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LocalReader reads files from the local filesystem
type LocalReader struct{}

// Open returns a ReadCloser for the specified target file
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

// List returns a list of file names in the local filesystem with the specified prefix
func (lr *LocalReader) List(prefix string) ([]string, error) {
	fileNames, err := filepath.Glob(fmt.Sprintf("%s/*.md", prefix))
	if err != nil {
		return nil, fmt.Errorf("error listing files: %w", err)
	}
	return fileNames, nil
}
