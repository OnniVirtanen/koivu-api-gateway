package server

import (
	"fmt"
	"koivu/gateway/auth"
	"koivu/gateway/config"
	"net/http"
	"net/url"
)

func Run() error {
	// Load the API key configuration
	authConfig, err := auth.LoadAuthConfiguration("../keys.yaml")
	if err != nil {
		return fmt.Errorf("could not load API key configuration: %v", err)
	}

	// Load the server configuration (not shown here, assume it's loaded similarly)
	serverConfig, err := config.LoadConfig("../config.yaml") // Assuming you have a function to load server config
	if err != nil {
		return fmt.Errorf("could not load server configuration: %v", err)
	}

	mux := http.NewServeMux()
	for _, route := range serverConfig.Routes {
		targetURL, err := url.Parse(route.Destination)
		if err != nil {
			return fmt.Errorf("invalid destination URL: %v", err)
		}

		proxy := NewProxy(targetURL)
		handler := ProxyRequestHandler(proxy, targetURL, route.Prefix)

		// Wrap the handler with http.HandlerFunc
		mux.Handle(route.Prefix, auth.AuthMiddleware(authConfig, route.Authentication, http.HandlerFunc(handler)))
	}

	address := "localhost:" + serverConfig.Port
	fmt.Println("Starting server on port", serverConfig.Port)
	if err := http.ListenAndServe(address, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
