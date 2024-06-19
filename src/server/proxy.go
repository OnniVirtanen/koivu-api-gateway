package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// NewProxy creates a reverse proxy for the given target URL.
func NewProxy(target url.URL) *httputil.ReverseProxy {
	target.Path = ""
	return httputil.NewSingleHostReverseProxy(&target)
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy, targetUrl *url.URL, routePrefix string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Original request path: %s\n", r.URL.Path)

		// Update the headers for the proxy
		r.URL = targetUrl
		r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

		// Serve the HTTP request through the proxy
		proxy.ServeHTTP(w, r)
	}
}
