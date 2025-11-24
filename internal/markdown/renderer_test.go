package markdown

import (
	"bytes"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/yuin/goldmark"
)

type stubReader struct {
	files map[string]string
}

func (s *stubReader) Open(target string) (io.ReadCloser, error) {
	c, ok := s.files[target]
	if !ok {
		return nil, ErrFileNotFound
	}
	return io.NopCloser(bytes.NewBufferString(c)), nil
}
func (s *stubReader) List(prefix string) ([]string, error) {
	var out []string
	for fn := range s.files {
		if filepath.Dir(fn) == prefix {
			out = append(out, fn)
		}
	}
	return out, nil
}

func TestRendererRender(t *testing.T) {
	cfg, _ := config.New()
	cfg.PostBasePath = "testdata"

	stub := &stubReader{files: map[string]string{
		"testdata/example.md": "---\n title: Example Title\n description: Desc\n image: img.png\n date: 2025-11-24T00:00:00Z\n---\n\n# Heading\n\nContent here.",
	}}
	r := &Renderer{Reader: stub, BasePath: cfg.PostBasePath, Markdown: testMarkdown()}
	content, meta, err := r.Render("example")
	if err != nil {
		t.Fatalf("render error: %v", err)
	}
	if meta.Slug != "example" {
		t.Errorf("expected slug example, got %s", meta.Slug)
	}
	if !strings.Contains(content, "<h1>") {
		t.Errorf("expected heading rendered, got %s", content)
	}
}

func TestRendererListSort(t *testing.T) {
	cfg, _ := config.New()
	cfg.PostBasePath = "testdata"
	stub := &stubReader{files: map[string]string{
		"testdata/a.md": "---\n title: A\n date: 2024-01-01T00:00:00Z\n---\n",
		"testdata/b.md": "---\n title: B\n date: 2025-01-01T00:00:00Z\n---\n",
	}}
	r := &Renderer{Reader: stub, BasePath: cfg.PostBasePath, Markdown: testMarkdown()}
	list, err := r.List()
	if err != nil {
		t.Fatalf("list error: %v", err)
	}
	if len(list) != 2 {
		t.Fatalf("expected 2 posts, got %d", len(list))
	}
	if list[0].Title != "B" {
		t.Errorf("expected first post B (newest), got %s", list[0].Title)
	}
}

// testMarkdown returns a minimal goldmark instance; using real one adds overhead.
func testMarkdown() goldmark.Markdown {
	return goldmark.New()
}
