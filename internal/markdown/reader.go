package markdown

import "io"

// FileReader t.b.d. until API stable
type FileReader interface {
	Open(target string) (io.ReadCloser, error)
	List(prefix string) ([]string, error)
}
