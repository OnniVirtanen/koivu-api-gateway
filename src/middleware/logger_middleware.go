package middleware

import (
	"io/ioutil"
	"koivu/gateway/logger"
	"net/http"
	"time"
)

// LoggerMiddleware logs details about the request and response
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var requestBody string
		if r.Body != nil {
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				r.Body = ioutil.NopCloser(r.Body)
			}
		}

		// Log request details
		logger.Log("Request: %s %s %s", r.Method, r.RequestURI, r.Proto)
		logger.Log("Remote Address: %s", r.RemoteAddr)
		logger.Log("User-Agent: %s", r.Header.Get("User-Agent"))
		logger.Log("Referer: %s", r.Header.Get("Referer"))
		logger.Log("Request Body: %s", requestBody)

		// Create a response writer to capture the status code
		rw := &responseWriter{w, http.StatusOK}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Log the response status and the time taken
		logger.Log("Response Status: %d", rw.statusCode)
		logger.Log("Request Duration: %s", time.Since(start))
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
