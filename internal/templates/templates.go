package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

var templates map[string]*template.Template

// LoadTemplates t.b.d. until API stable
func LoadTemplates() error {
	templates = make(map[string]*template.Template)

	partials, err := filepath.Glob("templates/partials/*.gohtml")
	if err != nil {
		return err
	}

	tmplFiles, err := filepath.Glob("templates/pages/*.gohtml")
	if err != nil {
		return err
	}

	for _, tmplFile := range tmplFiles {
		tmplName := strings.TrimSuffix(filepath.Base(tmplFile), ".gohtml")
		partials = append(partials, tmplFile)

		t, err := template.New("").Funcs(template.FuncMap{
			"mod": func(a, b int) int { return a % b },
		}).ParseFiles(partials...)
		if err != nil {
			return err
		}
		templates[tmplName] = t
	}

	return nil
}

// RenderTemplate t.b.d. until API stable
func RenderTemplate[T any](w http.ResponseWriter, tmpl string, data *models.PageData[T]) {
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
