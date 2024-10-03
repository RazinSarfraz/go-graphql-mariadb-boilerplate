package middleware

import (
	"net/http"
)

// Middleware function that prints a message before handling the request.
func DummyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print the message
		//log.Println("I am the middleware")
		// Call the next handler (GraphQL server in this case)
		next.ServeHTTP(w, r)
	})
}
