package server

import (
	"fmt"
	"koivu/gateway/config"
	"log"
	"net/http"
	"net/url"
)

func Run() error {
	config, err := config.LoadConfig("../config.yaml")
	if err != nil {
		fmt.Errorf("could not load configuration: %v", err)
	}
	mux := http.NewServeMux()
	for _, route := range config.Routes {
		url, _ := url.Parse(route.Destination)
		proxy := NewProxy(url)
		mux.HandleFunc(route.Prefix, ProxyRequestHandler(proxy, url, route.Prefix))
	}
	if err := http.ListenAndServe("localhost:"+config.Port, mux); err != nil {
		return fmt.Errorf("could not start the server: %v", err)
	}
	log.Println("Started server and running on port:", config.Port)
	return nil
}
