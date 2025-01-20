package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Title string
	Path  string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *PageData) {
	partialsDir := "ui/templates/partials/"

	partials, err := filepath.Glob(partialsDir + "*.tmpl.html")
	if err != nil {
		log.Fatal("Error loading partials: ", err)
	}

	tmplFiles := append(partials, "ui/templates/pages/"+tmpl+".tmpl.html")

	t, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		log.Fatal("Error parsing templates: ", err)
	}

	err = t.ExecuteTemplate(w, "base", data)
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

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func main() {
	fs := http.FileServer(http.Dir("./ui/static"))

	http.Handle("GET /static/", http.StripPrefix("/static", fs))
	http.HandleFunc("GET /", ServeHTTP)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
