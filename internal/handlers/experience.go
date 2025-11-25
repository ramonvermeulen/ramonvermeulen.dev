package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/data"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// ExperienceHandler handles the experience page with static position data
func ExperienceHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		d := &models.PageData[[]*models.Position]{
			Title:       "Experience",
			Path:        r.URL.Path,
			Content:     data.Positions,
			CdnURL:      cfg.CdnURL,
			Description: "Overview of my professional experience as a Software and Cloud Engineer, highlighting my expertise in cloud computing, software development, and DevOps practices.",
		}

		templates.RenderTemplate[[]*models.Position](w, "experience", d)
	}
}
