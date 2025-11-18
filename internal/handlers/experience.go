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
		Companies: []*models.Company{
			{
				Name: "Xebia",
				Positions: []*models.Position{
					{
						StartDate:   "Oct 2025",
						EndDate:     "Present",
						Title:       "Cloud Engineer",
						Description: "To be determined",
						IsCurrent:   true,
					},
					{
						StartDate:   "Dec 2021",
						EndDate:     "Oct 2025",
						Title:       "Data Engineer",
						Description: "To be determined",
						IsCurrent:   false},
				},
			},
			{
				Name: "ShoppingMinds",
				Positions: []*models.Position{
					{
						StartDate:   "Nov 2020",
						EndDate:     "Nov 2021",
						Title:       "Software Engineer",
						Description: "To be determined",
						IsCurrent:   false,
					},
				},
			},
			{
				Name: "Coffee IT",
				Positions: []*models.Position{
					{
						StartDate:   "Feb 2019",
						EndDate:     "Oct 2020",
						Title:       "Software Engineer (Mobile Apps)",
						Description: "To be determined",
						IsCurrent:   false,
					},
				},
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
