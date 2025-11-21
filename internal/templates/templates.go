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

	pageFiles, err := filepath.Glob("templates/pages/*.gohtml")
	if err != nil {
		return err
	}

	fm := buildFuncMap()

	for _, page := range pageFiles {
		name := strings.TrimSuffix(filepath.Base(page), ".gohtml")
		files := append(append([]string{}, partials...), page) // fresh slice each iteration
		parsed, err := template.New(name).Funcs(fm).ParseFiles(files...)
		if err != nil {
			return err
		}
		templates[name] = parsed
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

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
