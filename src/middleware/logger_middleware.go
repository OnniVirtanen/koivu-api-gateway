package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs details about the request and response
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log request details
		log.Printf("Request: %s %s %s\n", r.Method, r.RequestURI, r.Proto)
		log.Printf("Remote Address: %s\n", r.RemoteAddr)

		// If you want to log headers, uncomment the following lines
		for name, values := range r.Header {
			for _, value := range values {
				log.Printf("Header: %s: %s\n", name, value)
			}
		}

		// Create a response writer to capture the status code
		rw := &responseWriter{w, http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the response status and the time taken
		log.Printf("Response Status: %d\n", rw.statusCode)
		log.Printf("Request Duration: %s\n", time.Since(start))
	})
}

// responseWriter is a custom http.ResponseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code and calls the original WriteHeader
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
