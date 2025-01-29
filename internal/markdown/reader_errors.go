package markdown

import "errors"

var (
	ErrFileNotFound = errors.New("file not found")
	ErrReadFailed   = errors.New("failed to read file")
)
