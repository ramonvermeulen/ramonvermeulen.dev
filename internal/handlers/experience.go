package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// ExperienceHandler handles the experience page with static position data
func ExperienceHandler(cfg *config.Config) http.HandlerFunc {
	experienceData := &models.Experience{
		Positions: []*models.Position{
			{
				StartDate:   "Dec 2021",
				EndDate:     "Present",
				Title:       "Cloud Data Engineer",
				Company:     "Xebia",
				Description: "To be determined",
				IsCurrent:   true,
			},
			{
				StartDate:   "Nov 2020",
				EndDate:     "Nov 2021",
				Title:       "Software Engineer",
				Company:     "ShoppingMinds",
				Description: "To be determined",
				IsCurrent:   false,
			},
			{
				StartDate:   "Feb 2019",
				EndDate:     "Oct 2020",
				Title:       "Software Engineer (Mobile Apps)",
				Company:     "Coffee IT",
				Description: "To be determined",
				IsCurrent:   false,
			},
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		data := &models.PageData[models.Experience]{
			Title:   "Experience",
			Path:    r.URL.Path,
			Content: experienceData,
			CdnURL:  cfg.CdnURL,
		}

		templates.RenderTemplate[models.Experience](w, "experience", data)
	}
}
