package main

import (
	"bytes"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/yuin/goldmark"
	"io"
	"os"

	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templates map[string]*template.Template

type PageData struct {
	Title   string
	Path    string
	Content template.HTML
}

type FileReader interface {
	Read(path string) ([]byte, error)
}

type MarkdownReader struct{}

func (mr *MarkdownReader) Read(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *PageData) {
	t, ok := templates[tmpl]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	err := t.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Fatal(err)
	}
}

var routeMap = map[string]struct {
	template string
	title    string
}{
	"/":           {"about", "About"},
	"/experience": {"experience", "Experience"},
	"/blog":       {"blog", "Blog"},
}

func StaticPageHandler(w http.ResponseWriter, r *http.Request) {
	route, exists := routeMap[r.URL.Path]
	if !exists {
		http.NotFound(w, r)
		return
	}

	data := &PageData{
		Title: route.title,
		Path:  r.URL.Path,
	}

	renderTemplate(w, route.template, data)
}

func BlogPostHandler(fileReader FileReader) http.HandlerFunc {
	renderer := goldmark.New()
	return func(w http.ResponseWriter, r *http.Request) {
		postSlug := chi.URLParam(r, "postSlug")
		markdownPath := fmt.Sprintf("./posts/%s.md", postSlug)
		markdown, err := fileReader.Read(markdownPath)
		if err != nil {
			/* handle different errors */
			http.NotFound(w, r)
			return
		}
		var buf bytes.Buffer
		err = renderer.Convert(markdown, &buf)
		if err != nil {
			/* handle different errors */
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}

		data := &PageData{
			Title:   postSlug,
			Path:    r.URL.Path,
			Content: template.HTML(buf.String()),
		}

		renderTemplate(w, "post", data)
	}
}

func loadTemplates() error {
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
		tmplName := filepath.Base(tmplFile)
		tmplName = tmplName[:len(tmplName)-len(".tmpl.html")] // Remove the file extension

		files := append(partials, tmplFile)
		t, err := template.ParseFiles(files...)
		if err != nil {
			return err
		}
		templates[tmplName] = t
	}

	return nil
}

func main() {
	if err := loadTemplates(); err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	router := chi.NewRouter()
	fs := http.FileServer(http.Dir("./assets/static"))

	router.Handle("GET /static/*", http.StripPrefix("/static", fs))
	router.Get("/blog/{postSlug:[a-z-]+}", BlogPostHandler(&MarkdownReader{}))
	router.Get("/*", StaticPageHandler)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
