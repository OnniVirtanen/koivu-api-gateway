package server

import (
	"fmt"
	"koivu/gateway/config"
	"koivu/gateway/middleware"
	"net/http"
	"net/url"
)

func Run() error {
	config.InitConfig()
	config := *config.GetConfig()

	mux := http.NewServeMux()
	for _, route := range config.RouteConfiguration.Routes {
		fmt.Println("route: ", route)
		targetURL, err := url.Parse(route.Destination)
		if err != nil {
			return fmt.Errorf("invalid destination URL: %v", err)
		}

		proxy := NewProxy(targetURL)
		handler := ProxyRequestHandler(proxy, targetURL, route.Prefix)

		finalHandler := middleware.AuthMiddleware(config.AuthConfiguration, route.Authentication, http.HandlerFunc(handler))
		finalHandler = middleware.RateLimitMiddleware(&route.RateLimitConfiguration, finalHandler)

		mux.Handle(route.Prefix, finalHandler)
	}

	address := "localhost:" + config.RouteConfiguration.Port
	fmt.Println("Starting server on port", config.RouteConfiguration.Port)
	if err := http.ListenAndServe(address, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}
