package main

import (
	"koivu/gateway/server"
	"log"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
