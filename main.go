package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/handlers"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}
	if err := templates.LoadTemplates(); err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	renderer, err := markdown.NewRenderer(cfg)
	if err != nil {
		log.Fatal("Error creating renderer: ", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	if cfg.Env == "dev" {
		templates.SetDevMode(true)
		fs := http.FileServer(http.Dir("./public"))
		router.Handle("/public/*", http.StripPrefix("/public", fs))
	}

	router.Get("/robots.txt", handlers.RobotsTxtHandler(cfg))
	router.Get("/sitemap.xml", handlers.SitemapXMLHandler(cfg, renderer))
	router.Get("/ping", handlers.PongHandler())
	router.Get("/blog", handlers.BlogIndexHandler(cfg, renderer))
	router.Get("/blog/{postSlug:[a-z-]+}", handlers.BlogPostHandler(cfg, renderer))
	router.Get("/experience", handlers.ExperienceHandler(cfg))
	router.Get("/*", handlers.StaticPageHandler(cfg))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
