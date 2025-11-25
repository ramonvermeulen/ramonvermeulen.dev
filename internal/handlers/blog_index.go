package handlers

import (
	"log"
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// BlogIndexHandler renders the blog index page.
func BlogIndexHandler(cfg *config.Config, renderer *markdown.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := renderer.List()
		if err != nil {
			log.Printf("Error listing posts: %v", err)
			http.Error(w, "Error listing posts", http.StatusInternalServerError)
			return
		}

		data := &models.PageData[models.BlogIndex]{
			Title:       "Blog",
			Description: "Collection of personal articles on software development, cloud computing, and technology.",
			Path:        r.URL.Path,
			CdnURL:      cfg.CdnURL,
			Content: models.BlogIndex{
				Posts: posts,
			},
		}

		templates.RenderTemplate(w, "blog", data)
	}
}
