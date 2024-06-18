package main

import (
	"fmt"
	"koivu/gateway/config"
	"log"
)

func main() {
	config, err := config.LoadConfig("../config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	fmt.Printf("Port: %d\n", config.Port)
	for _, route := range config.Routes {
		fmt.Printf("Route Prefix: %s, Destination: %s\n", route.Prefix, route.Destination)
	}
}
