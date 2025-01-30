package markdown

import "errors"

var (
	// ErrFileNotFound t.b.d. until API stable
	ErrFileNotFound = errors.New("file not found")
	// ErrReadFailed t.b.d. until API stable
	ErrReadFailed = errors.New("failed to read file")
)
