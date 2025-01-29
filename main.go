package main

import (
	"context"
	"log"
	"net/http"

	"errors"
	"os"

	"cloud.google.com/go/storage"
	"github.com/go-chi/chi/v5"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/handlers"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/templates"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting/v2"
)

func newFileReader() (markdown.FileReader, error) {
	basePath := os.Getenv("POSTS_BASE_PATH")
	if basePath == "" {
		log.Println("POSTS_BASE_PATH environment variable not set, defaulting to 'posts'")
		basePath = "./posts"
	}
	if os.Getenv("ENV") == "production" {
		bucketName := os.Getenv("GCS_POSTS_BUCKET")
		if bucketName == "" {
			return nil, errors.New("GCS_POSTS_BUCKET environment variable is required in production")
		}
		ctx := context.Background()
		client, err := storage.NewClient(ctx)
		if err != nil {
			return nil, err
		}
		return &markdown.GCSReader{
			BucketName: bucketName,
			Client:     client,
			BasePath:   basePath,
		}, nil
	}
	return &markdown.LocalReader{
		BasePath: basePath,
	}, nil
}

func main() {
	if err := templates.LoadTemplates(); err != nil {
		log.Fatal("Error loading templates: ", err)
	}

	reader, err := newFileReader()
	if err != nil {
		log.Fatal("Error creating reader: ", err)
	}

	renderer := &markdown.Renderer{
		Reader: reader,
		Markdown: goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("nord"),
				),
			),
		),
	}

	router := chi.NewRouter()
	fs := http.FileServer(http.Dir("./assets/static"))

	router.Handle("GET /static/*", http.StripPrefix("/static", fs))
	router.Get("/blog/{postSlug:[a-z-]+}", handlers.BlogPostHandler(renderer))
	router.Get("/*", handlers.StaticPageHandler())

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
