package markdown

import (
	"fmt"
	"io"
	"os"
)

type LocalReader struct {
	BasePath string
}

func (lr *LocalReader) Read(target string) ([]byte, error) {
	markdownPath := fmt.Sprintf("%s/%s.md", lr.BasePath, target)
	file, err := os.Open(markdownPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		}
		return nil, ErrReadFailed
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, ErrReadFailed
	}
	return content, nil
}
