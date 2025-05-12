package markdown

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

// Renderer t.b.d. until API stable
type Renderer struct {
	Reader   FileReader
	Markdown goldmark.Markdown
	BasePath string
}

// Render t.b.d. until API stable
func (r *Renderer) Render(target string) (string, *models.BlogPostMeta, error) {
	target = fmt.Sprintf("%s/%s.md", r.BasePath, target)
	reader, err := r.Reader.Open(target)
	if err != nil {
		return "", nil, fmt.Errorf("failed to open file: %w", err)
	}

	var meta models.BlogPostMeta
	remainder, err := frontmatter.Parse(reader, &meta)
	if err != nil {
		return "", nil, fmt.Errorf("failed to parse frontmatter: %w", err)
	}

	var buf bytes.Buffer
	err = r.Markdown.Convert(remainder, &buf)
	if err != nil {
		return "", nil, fmt.Errorf("failed to convert markdown: %w", err)
	}

	return buf.String(), &meta, nil
}

// List t.b.d. until API stable
func (r *Renderer) List() ([]*models.BlogPostMeta, error) {
	var posts []*models.BlogPostMeta
	fileNames, err := r.Reader.List(r.BasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %w", err)
	}
	for _, fileName := range fileNames {
		reader, err := r.Reader.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("failed to open file: %w", err)
		}
		var meta models.BlogPostMeta
		_, err = frontmatter.Parse(reader, &meta)
		if err != nil {
			return nil, fmt.Errorf("failed to parse frontmatter: %w", err)
		}
		meta.Slug = strings.ToLower(strings.TrimSuffix(filepath.Base(fileName), ".md"))
		posts = append(posts, &meta)
	}
	return posts, nil
}

// NewRenderer t.b.d. until API stable
func NewRenderer(cfg *config.Config) (*Renderer, error) {
	reader, err := NewFileReader(cfg)
	if err != nil {
		return nil, err
	}

	return &Renderer{
		Reader: reader,
		Markdown: goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("nord"),
				),
			),
		),
		BasePath: cfg.PostBasePath,
	}, nil
}
