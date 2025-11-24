package markdown

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"path/filepath"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

// Renderer renders markdown files to HTML with metadata.
type Renderer struct {
	Reader   FileReader
	Markdown goldmark.Markdown
	BasePath string
}

// Render renders the markdown file identified by the slug to HTML and extracts its metadata.
func (r *Renderer) Render(slug string) (string, *models.BlogPostMeta, error) {
	slug = strings.ToLower(strings.TrimSpace(slug))
	if slug == "" {
		return "", nil, fmt.Errorf("empty slug")
	}
	target := filepath.Join(r.BasePath, slug+".md")

	reader, err := r.Reader.Open(target)
	if err != nil {
		return "", nil, fmt.Errorf("open %s: %w", target, err)
	}
	defer func() {
		if cerr := reader.Close(); cerr != nil {
			log.Printf("error closing reader for %s: %v", target, cerr)
		}
	}()

	var meta models.BlogPostMeta
	remainder, err := frontmatter.Parse(reader, &meta)
	if err != nil {
		return "", nil, fmt.Errorf("parse frontmatter %s: %w", target, err)
	}

	var buf bytes.Buffer
	if err = r.Markdown.Convert(remainder, &buf); err != nil {
		return "", nil, fmt.Errorf("convert markdown %s: %w", target, err)
	}

	meta.Slug = slug
	return buf.String(), &meta, nil
}

// List lists all markdown files in the base path and returns their metadata sorted by date descending.
func (r *Renderer) List() ([]*models.BlogPostMeta, error) {
	fileNames, err := r.Reader.List(r.BasePath)
	if err != nil {
		return nil, fmt.Errorf("list posts: %w", err)
	}
	posts := make([]*models.BlogPostMeta, 0, len(fileNames))
	for _, fileName := range fileNames {
		reader, err := r.Reader.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("open %s: %w", fileName, err)
		}
		content, err := io.ReadAll(reader)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", fileName, err)
		}
		if cerr := reader.Close(); cerr != nil {
			log.Printf("error closing reader for %s: %v", fileName, cerr)
		}
		var meta models.BlogPostMeta
		if _, err = frontmatter.Parse(bytes.NewReader(content), &meta); err != nil {
			return nil, fmt.Errorf("frontmatter %s: %w", fileName, err)
		}
		meta.Slug = strings.ToLower(strings.TrimSuffix(filepath.Base(fileName), ".md"))
		posts = append(posts, &meta)
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	return posts, nil
}

// NewRenderer creates a new Renderer with the given configuration.
func NewRenderer(cfg *config.Config) (*Renderer, error) {
	reader, err := NewFileReader(cfg)
	if err != nil {
		return nil, fmt.Errorf("init reader: %w", err)
	}
	md := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(highlighting.WithStyle("nord")),
		),
	)
	return &Renderer{
		Reader:   reader,
		Markdown: md,
		BasePath: cfg.PostBasePath,
	}, nil
}
