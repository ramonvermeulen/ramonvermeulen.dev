package templates

import (
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

func TestLoadTemplatesAndRender(t *testing.T) {
	// Ensure test runs from repository root for relative glob patterns
	if _, err := os.Stat("templates/partials/base.gohtml"); err != nil {
		t.Skip("skipping: templates directory not found in CWD")
	}
	if err := LoadTemplates(); err != nil {
		t.Fatalf("load templates: %v", err)
	}
	w := httptest.NewRecorder()
	data := &models.PageData[models.NoContent]{
		Title:        "About",
		Path:         "/",
		CanonicalURL: "https://ramonvermeulen.dev/",
		Content:      models.NoContent{},
		CdnURL:       "http://cdn",
	}
	RenderTemplate[models.NoContent](w, "about", data)
	out := w.Body.String()
	if !strings.Contains(out, "<title>About | ramonvermeulen.dev</title>") {
		t.Errorf("expected title tag in output")
	}
	if !strings.Contains(out, `<link rel="canonical" href="https://ramonvermeulen.dev/"/>`) {
		t.Errorf("expected canonical tag in output")
	}
	if !strings.Contains(out, "Software & Cloud Engineer") {
		t.Errorf("expected snippet from about page; output length %d", len(out))
	}
}
