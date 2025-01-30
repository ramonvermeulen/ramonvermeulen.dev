package markdown

import (
	"io"
	"os"
)

// LocalReader t.b.d. until API stable
type LocalReader struct{}

// Read t.b.d. until API stable
func (lr *LocalReader) Read(target string) ([]byte, error) {
	file, err := os.Open(target)
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
