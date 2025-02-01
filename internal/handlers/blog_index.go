package handlers

import (
	"net/http"

	"github.com/ramonvermeulen/ramonvermeulen.dev/internal/markdown"
)

// BlogIndexHandler t.b.d. until API stable
func BlogIndexHandler(renderer *markdown.Renderer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//			posts, err := renderer.List()
	}
}
