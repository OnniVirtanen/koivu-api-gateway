package middleware

import (
	"koivu/gateway/config"
	"net/http"
	"strings"
)

func AuthMiddleware(masterConfiguration *config.Configuration, authType config.AuthType, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch authType {
		case config.Key:
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				http.Error(w, "API key required", http.StatusUnauthorized)
				return
			}
			if !isValidAPIKey(*masterConfiguration.AuthConfiguration, apiKey, getRouteNameByPath(r.URL.Path, masterConfiguration.RouteConfiguration)) {
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

func getRouteNameByPath(path string, config *config.RouteConfiguration) string {
	for _, route := range config.Routes {
		if strings.HasPrefix(path, route.Prefix) {
			return route.Name
		}
	}
	return ""
}

func isValidAPIKey(authConfig config.AuthConfiguration, apiKey string, routeName string) bool {
	for _, key := range authConfig.APIKeys {
		if key.Value == apiKey {

			for _, route := range key.Routes {
				if route == routeName {
					return true
				}
			}
		}
	}
	return false
}
