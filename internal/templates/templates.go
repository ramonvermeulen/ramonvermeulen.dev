package templates

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

var (
	templates map[string]*template.Template
	devMode   bool
	mu        sync.RWMutex // protects templates map
)

// buildFuncMap constructs the FuncMap used in templates.
func buildFuncMap() template.FuncMap {
	return template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"mod": func(a, b int) int { return a % b },
	}
}

// SetDevMode enables template reloading on each render in development.
func SetDevMode(enabled bool) { devMode = enabled }

// LoadTemplates parses all page + partial templates into memory.
func LoadTemplates() error {
	partials, err := filepath.Glob("templates/partials/*.gohtml")
	if err != nil {
		return err
	}
	pageFiles, err := filepath.Glob("templates/pages/*.gohtml")
	if err != nil {
		return err
	}

	fm := buildFuncMap()
	loaded := make(map[string]*template.Template, len(pageFiles))
	for _, page := range pageFiles {
		name := strings.TrimSuffix(filepath.Base(page), ".gohtml")
		files := append(append([]string{}, partials...), page) // fresh slice each iteration
		parsed, err := template.New(name).Funcs(fm).ParseFiles(files...)
		if err != nil {
			return err
		}
		loaded[name] = parsed
	}

	mu.Lock()
	templates = loaded
	mu.Unlock()
	return nil
}

// RenderTemplate executes the named template against base layout.
// It writes generic error messages to the client but logs detailed ones.
func RenderTemplate[T any](w http.ResponseWriter, tmpl string, data *models.PageData[T]) {
	if devMode { // hot reload in dev
		if err := LoadTemplates(); err != nil {
			log.Printf("template reload failed: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	mu.RLock()
	t, ok := templates[tmpl]
	mu.RUnlock()
	if !ok {
		log.Printf("template '%s' not found", tmpl)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := t.ExecuteTemplate(w, "base", data); err != nil {
		log.Printf("render '%s' failed: %v", tmpl, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
