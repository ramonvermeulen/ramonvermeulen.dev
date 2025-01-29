package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

var templates map[string]*template.Template

func LoadTemplates() error {
	templates = make(map[string]*template.Template)

	partials, err := filepath.Glob("templates/partials/*.tmpl.html")
	if err != nil {
		return err
	}

	tmplFiles, err := filepath.Glob("templates/pages/*.tmpl.html")
	if err != nil {
		return err
	}

	for _, tmplFile := range tmplFiles {
		tmplName := strings.TrimSuffix(filepath.Base(tmplFile), ".tmpl.html")
		files := append(partials, tmplFile)

		t, err := template.ParseFiles(files...)
		if err != nil {
			return err
		}
		templates[tmplName] = t
	}

	return nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.PageData) {
	t, ok := templates[tmpl]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
