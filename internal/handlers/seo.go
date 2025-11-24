package handlers

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/models"
)

// RobotsTxtHandler serves the /robots.txt file.
func RobotsTxtHandler(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		sitemapURL := fmt.Sprintf("%s/sitemap.xml", cfg.BaseURL)
		content := fmt.Sprintf("User-agent: *\nAllow: /\nSitemap: %s", sitemapURL)
		if _, err := w.Write([]byte(content)); err != nil {
			log.Printf("error writing robots.txt: %v", err)
		}
	}
}

// SitemapXMLHandler generates and serves the /sitemap.xml file.
func SitemapXMLHandler(cfg *config.Config, renderer *markdown.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		urls := []models.SitemapURL{
			{Loc: cfg.BaseURL, ChangeFreq: "weekly", Priority: 1.0},
			{Loc: fmt.Sprintf("%s/experience", cfg.BaseURL), ChangeFreq: "monthly", Priority: 1.0},
			{Loc: fmt.Sprintf("%s/blog", cfg.BaseURL), ChangeFreq: "weekly", Priority: 0.9},
		}

		postMeta, err := renderer.List()
		if err != nil {
			log.Printf("Error listing posts for sitemap: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for _, meta := range postMeta {
			urls = append(urls, models.SitemapURL{
				Loc:        fmt.Sprintf("%s/blog/%s", cfg.BaseURL, meta.Slug),
				LastMod:    meta.Date.Format("2006-01-02"),
				ChangeFreq: "yearly",
				Priority:   0.7,
			})
		}

		urlSet := models.URLSet{
			Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
			URLs:  urls,
		}

		w.Header().Set("Content-Type", "application/xml")
		if _, err := w.Write([]byte(xml.Header)); err != nil {
			log.Printf("error writing sitemap.xml header: %v", err)
			return
		}
		encoder := xml.NewEncoder(w)
		encoder.Indent("", "  ")
		if err := encoder.Encode(urlSet); err != nil {
			log.Printf("Error encoding sitemap: %v", err)
		}
	}
}
