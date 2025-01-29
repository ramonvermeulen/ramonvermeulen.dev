package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"

	"github.com/go-chi/chi/v5"
)

func StaticPageHandler() http.HandlerFunc {
	routeMap := map[string]struct {
		template string
		title    string
	}{
		"/":           {"about", "About"},
		"/experience": {"experience", "Experience"},
		"/blog":       {"blog", "Blog"},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		route, exists := routeMap[r.URL.Path]
		if !exists {
			http.NotFound(w, r)
			return
		}

		data := &models.PageData{
			Title: route.title,
			Path:  r.URL.Path,
		}

		templates.RenderTemplate(w, route.template, data)
	}
}

func BlogPostHandler(renderer *markdown.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postSlug := chi.URLParam(r, "postSlug")
		rendered, err := renderer.Render(postSlug)

		if err != nil {
			if errors.Is(err, markdown.ErrFileNotFound) {
				http.NotFound(w, r)
			} else if errors.Is(err, markdown.ErrReadFailed) {
				http.Error(w, "Error reading or converting markdown", http.StatusInternalServerError)
			} else {
				http.Error(w, "Error rendering post", http.StatusInternalServerError)
			}
			return
		}

		data := &models.PageData{
			Title:   postSlug,
			Path:    r.URL.Path,
			Content: template.HTML(rendered),
		}

		templates.RenderTemplate(w, "post", data)
	}
}
