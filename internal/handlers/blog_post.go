package handlers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// BlogPostHandler renders a single blog post based on the postSlug URL parameter.
func BlogPostHandler(cfg *config.Config, renderer *markdown.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postSlug := chi.URLParam(r, "postSlug")
		rendered, meta, err := renderer.Render(postSlug)

		if err != nil {
			switch {
			case errors.Is(err, markdown.ErrFileNotFound):
				http.NotFound(w, r)
			case errors.Is(err, markdown.ErrReadFailed):
				http.Error(w, "Error reading or converting markdown", http.StatusInternalServerError)
			default:
				http.Error(w, "Error rendering post", http.StatusInternalServerError)
			}
			return
		}

		data := &models.PageData[models.BlogPost]{
			Title:       meta.Title,
			Description: meta.Description,
			Path:        r.URL.Path,
			CdnURL:      cfg.CdnURL,
			Content: models.BlogPost{
				Content: template.HTML(rendered),
				Meta:    meta,
			},
		}

		templates.RenderTemplate[models.BlogPost](w, "post", data)
	}
}
