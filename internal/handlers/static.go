package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// StaticPageHandler handles static pages like the about page.
func StaticPageHandler(cfg *config.Config) http.HandlerFunc {
	routeMap := map[string]struct {
		template    string
		title       string
		description string
	}{
		"/": {"about", "About", "I am a Software and Cloud Engineer with a strong background in Google Cloud Platform, currently working at Xebia. My journey at Xebia began as a Data Engineer, and my interests have steadily evolved towards Cloud Infrastructure and Software Engineering."},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		route, exists := routeMap[r.URL.Path]
		if !exists {
			http.NotFound(w, r)
			return
		}

		data := &models.PageData[models.NoContent]{
			Title:   route.title,
			Path:    r.URL.Path,
			Content: models.NoContent{},
			CdnURL:  cfg.CdnURL,
		}

		templates.RenderTemplate[models.NoContent](w, route.template, data)
	}
}
