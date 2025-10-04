package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log method and path
		log.Printf("%s %s", r.Method, r.URL.Path)

		// Check for X-API-Key header
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret123" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "unauthorized"}`)
			return
		}

		// If valid, proceed to next handler
		next.ServeHTTP(w, r)
	})
}
