package middleware

import (
	"koivu/gateway/config"
	"net/http"
)

func AuthMiddleware(authConfig *config.AuthConfiguration, authType config.AuthType, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch authType {
		case config.Key:
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				http.Error(w, "API key required", http.StatusUnauthorized)
				return
			}
			if !isValidAPIKey(*authConfig, apiKey, r.URL.Path) {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}
		case config.None:
			// No authentication required
		default:
			http.Error(w, "Unsupported authentication type", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isValidAPIKey(authConfig config.AuthConfiguration, apiKey string, path string) bool {
	for _, key := range authConfig.APIKeys {
		if key.Value == apiKey {

			for _, route := range key.Routes {
				if route == path {
					return true
				}
			}
		}
	}
	return false
}
