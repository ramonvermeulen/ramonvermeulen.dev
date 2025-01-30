package markdown

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/yuin/goldmark"
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
	log.Print(target)
	content, err := r.Reader.Read(target)
	if err != nil {
		return "", nil, err
	}

	var meta models.BlogPostMeta
	// TODO: refactor r.Reader.Read() in some way so that frontmatter.Parse()
	// doesn't have to make a reader from the content again
	remainder, err := frontmatter.Parse(strings.NewReader(string(content)), &meta)
	if err != nil {
		return "", nil, err
	}

	var buf bytes.Buffer
	err = r.Markdown.Convert(remainder, &buf)
	if err != nil {
		return "", nil, err
	}

	return buf.String(), nil, nil
}
