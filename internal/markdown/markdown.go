package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
)

type Renderer struct {
	Reader   FileReader
	Markdown goldmark.Markdown
}

func (r *Renderer) Render(target string) (string, error) {
	content, err := r.Reader.Read(target)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = r.Markdown.Convert(content, &buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
