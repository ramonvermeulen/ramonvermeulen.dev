package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/config"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/handlers"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func newFileReader(cfg *config.Config) (markdown.FileReader, error) {
	if cfg.Env != "dev" {
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		return &markdown.GCSReader{
			BucketName: cfg.GCSBucket,
			Client:     client,
		}, nil
	}
	return &markdown.LocalReader{}, nil
}

func newRenderer(cfg *config.Config) (*markdown.Renderer, error) {
	reader, err := newFileReader(cfg)
	if err != nil {
		return nil, err
	}

	return &markdown.Renderer{
		Reader: reader,
		Markdown: goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("nord"),
				),
			),
		),
		BasePath: cfg.PostBasePath,
	}, nil
}

func main() {
	cfg := config.New()
	if err := templates.LoadTemplates(); err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	renderer, err := newRenderer(cfg)
	if err != nil {
		log.Fatal("Error creating renderer: ", err)
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	if cfg.Env == "dev" {
		fs := http.FileServer(http.Dir("./public"))
		router.Handle("GET /public/*", http.StripPrefix("/public", fs))
	}

	router.Get("/blog", handlers.BlogIndexHandler(cfg, renderer))
	router.Get("/blog/{postSlug:[a-z-]+}", handlers.BlogPostHandler(cfg, renderer))
	router.Get("/*", handlers.StaticPageHandler(cfg))
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
