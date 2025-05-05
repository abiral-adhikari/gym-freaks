package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func GetTokenFromRequest(r *http.Request) (string, error) {
	// Get the Authorization header from the request
	authHeader := r.Header.Get("Authorization")

	// Check if the header exists and starts with "Bearer "
	if authHeader == "" {
		return "", fmt.Errorf("authorization header missing")
	}

	// Split "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid Authorization header format")
	}

	// Return the token part
	return parts[1], nil
}



func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for specific routes (e.g., /signup, /login, etc.)
		// if strings.Contains(r.URL.Path, "/signup") || strings.Contains(r.URL.Path, "/login") {
		// 	next.ServeHTTP(w, r) // Proceed without authentication for these routes
		// 	return
		// }

		// Example of checking an Authorization token (for illustration purposes)
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		// Simulate token validation (for illustration purposes)
		if token != "Bearer valid_token" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Continue to the next handler if authentication is successful
		next.ServeHTTP(w, r)
	})
}
