package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

// ExperienceHandler handles the experience page with static position data
func ExperienceHandler(cfg *config.Config) http.HandlerFunc {
	positions := []*models.Position{
		{
			CompanyName:    "Xebia",
			CompanyWebsite: "https://xebia.com/",
			StartDate:      "Oct 2025",
			EndDate:        "Present",
			Title:          "Cloud Engineer",
			Location:       "Hilversum, the Netherlands",
			Description:    "Cloud Engineering consultant for the cloud branch of Xebia in the Netherlands (Xebia Cloud). Helping organizations build maintainable and scalable cloud native software solutions and infrastructure in the cloud.",
			IsCurrent:      true,
		},
		{
			CompanyName:    "Xebia",
			CompanyWebsite: "https://xebia.com/",
			StartDate:      "Dec 2021",
			EndDate:        "Oct 2025",
			Title:          "Data Engineer",
			Location:       "Amsterdam, the Netherlands",
			Description:    "Data Engineering consultant for the data branch of Xebia in the Netherlands (Xebia Data). Helping organizations build maintainable and scalable data solutions in the cloud.",
			IsCurrent:      false,
		},
		{
			CompanyName:    "ShoppingMinds",
			CompanyWebsite: "https://shoppingminds.com/",
			StartDate:      "Nov 2020",
			EndDate:        "Nov 2021",
			Title:          "Software Engineer",
			Location:       "Utrecht, the Netherlands",
			Description:    "Developed different features for Shopping Minds their data management platform, mainly in Python. Responsible for designing and developing connectors to external APIs, for instance to trigger Email campaigns, update audiences, or sync profile information.",
			IsCurrent:      false,
		},
		{
			CompanyName:    "Coffee IT",
			CompanyWebsite: "https://coffeeit.nl/",
			StartDate:      "Feb 2019",
			EndDate:        "Oct 2020",
			Title:          "Software Engineer (Mobile Apps)",
			Location:       "Utrecht, the Netherlands",
			Description:    "Worked as a mobile engineer and helped building several mobile applications (iOS and Android) for Coffee IT their clients (Knaek, ROC, Wiebetaaltwat, Examenoverzicht). Gained a lot of front-end experience, mainly in React (Native).",
			IsCurrent:      false,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		data := &models.PageData[[]*models.Position]{
			Title:   "Experience",
			Path:    r.URL.Path,
			Content: positions,
			CdnURL:  cfg.CdnURL,
		}

		templates.RenderTemplate[[]*models.Position](w, "experience", data)
	}
}
