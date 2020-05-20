package middleware

import (
	"log"
	"net/http"
	"strings"
)

func logRequest(details ...string) {
	log.Print(strings.Join(details, " - "))
}

// LogRequestHandler an Http Request
func LogRequestHandler(h http.Handler) http.Handler {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		// Call the original request
		h.ServeHTTP(w, r)

		// Log info here
		uri := r.URL.String()
		method := r.Method
		logRequest(uri, method)
	}

	return http.HandlerFunc(wrapped)
}
