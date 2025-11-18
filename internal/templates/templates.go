package templates

import (
	"html/template"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

var (
	templates map[string]*template.Template
	devMode   bool
)

// SetDevMode enables template reloading on each render in development.
func SetDevMode(enabled bool) { devMode = enabled }

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
			"sub": func(a, b int) int { return a - b },
			"mod": func(a, b int) int { return a % b },
			"add": func(a, b int) int { return a + b },
			"len": func(v interface{}) int {
				switch val := v.(type) {
				case []*models.Position:
					return len(val)
				default:
					return 0
				}
			},
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
	if devMode {
		if err := LoadTemplates(); err != nil {
			http.Error(w, "Error reloading templates", http.StatusInternalServerError)
			return
		}
	}

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
