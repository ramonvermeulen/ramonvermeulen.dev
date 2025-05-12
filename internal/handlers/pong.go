package handlers

import (
	"log"
	"net/http"
)

// PongHandler is a simple handler that responds with "pong" used for health checks.
func PongHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	}
}
