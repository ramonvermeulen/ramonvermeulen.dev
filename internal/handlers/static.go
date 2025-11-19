package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// StaticPageHandler t.b.d. until API stable
func StaticPageHandler(cfg *config.Config) http.HandlerFunc {
	routeMap := map[string]struct {
		template string
		title    string
	}{
		"/": {"about", "About"},
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
