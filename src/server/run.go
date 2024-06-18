package server

import (
	"fmt"
	"koivu/gateway/config"
	"koivu/gateway/middleware"
	"net/http"
	"net/url"
	"time"

	"github.com/common-nighthawk/go-figure"
)

const separator = "---------------------------------------------------"

func Run() error {
	startTime := time.Now()

	printASCII()
	config.InitConfig()
	config := *config.GetConfig()

	mux := http.NewServeMux()
	for _, route := range config.RouteConfiguration.Routes {
		targetURL, err := url.Parse(route.Destination)
		if err != nil {
			return fmt.Errorf("invalid destination URL: %v", err)
		}

		proxy := NewProxy(targetURL)
		handler := ProxyRequestHandler(proxy, targetURL, route.Prefix)

		finalHandler := middleware.AuthMiddleware(config.AuthConfiguration, route.Authentication, http.HandlerFunc(handler))
		finalHandler = middleware.RateLimitMiddleware(&route.RateLimitConfiguration, finalHandler)
		finalHandler = middleware.LoggerMiddleware(finalHandler)

		mux.Handle(route.Prefix, finalHandler)
	}

	printRoutes(*config.RouteConfiguration)
	fmt.Println(separator)
	fmt.Printf("Routes found %v\n", len(config.RouteConfiguration.Routes))
	fmt.Printf("API-keys found %v\n", len(config.AuthConfiguration.APIKeys))

	address := "localhost:" + config.RouteConfiguration.Port

	elapsedTime := time.Since(startTime)
	fmt.Println(separator)
	fmt.Printf("Startup time: %v\n", elapsedTime)
	fmt.Println("Listening on port: ", config.RouteConfiguration.Port)
	fmt.Println(separator)

	if err := http.ListenAndServe(address, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	return nil
}

func printASCII() {
	fmt.Println(separator)
	myFigure := figure.NewFigure("KOIVU-API-GATEWAY", "", true)
	myFigure.Print()
	fmt.Println(separator)
}

func printRoutes(config config.RouteConfiguration) {
	fmt.Println("API Gateway Routes")
	fmt.Printf("%-30s %-50s %-70s %-10s %-20s\n", "Name", "Prefix", "Destination", "Auth", "Rate Limit")
	fmt.Println(separator)

	for _, route := range config.Routes {
		rateLimit := fmt.Sprintf("%d:%s:%s", route.RateLimitConfiguration.Requests, route.RateLimitConfiguration.Timeframe, route.RateLimitConfiguration.Type)
		fmt.Printf("%-30s %-50s %-70s %-10s %-20s\n", truncateString(route.Name, 30), truncateString(route.Prefix, 50), truncateString(route.Destination, 70), truncateString(string(route.Authentication), 10), truncateString(rateLimit, 20))
	}
}

func truncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength-3] + "..."
	}
	return s
}
