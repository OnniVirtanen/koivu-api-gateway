package auth

import "net/http"

// AuthType represents the type of authentication used
type AuthType string

const (
	Key AuthType = "key"

	// None represents no authentication required
	None AuthType = "none"
)

func (a AuthType) IsAuthEnabled() bool {
	switch a {
	case None:
		return false
	default:
		return true
	}
}

func AuthMiddleware(authConfig *AuthConfiguration, authType AuthType, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch authType {
		case Key:
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				http.Error(w, "API key required", http.StatusUnauthorized)
				return
			}
			if !authConfig.IsValidAPIKey(apiKey, r.URL.Path) {
				http.Error(w, "Invalid API key", http.StatusUnauthorized)
				return
			}
		case None:
			// No authentication required
		default:
			http.Error(w, "Unsupported authentication type", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
