package main

import (
	"log"
	"net/http"

	"go-practice2/internal/handlers"
	"go-practice2/middleware"
)

func main() {
	// Create router
	mux := http.NewServeMux()

	// Apply middleware to both routes
	protectedHandler := middleware.AuthMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				handlers.GetUserHandler(w, r)
			case http.MethodPost:
				handlers.CreateUserHandler(w, r)
			default:
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write([]byte(`{"error": "method not allowed"}`))
			}
		}),
	)

	// Register the protected route
	mux.Handle("/user", protectedHandler)

	// Start server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
