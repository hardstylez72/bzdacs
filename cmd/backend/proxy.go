package main

import (
	"go.opencensus.io/plugin/ochttp"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func startProxyServer(apiServerHost, staticServerHost, proxyServerPort string, done chan struct{}) error {

	apiServerUrl, err := url.Parse(apiServerHost)
	if err != nil {
		return err
	}

	staticServerUrl, err := url.Parse(staticServerHost)
	if err != nil {
		return err
	}

	apiServerReverseProxy := httputil.NewSingleHostReverseProxy(apiServerUrl)
	apiServerReverseProxy.Transport = &ochttp.Transport{}

	staticServerReverseProxy := httputil.NewSingleHostReverseProxy(staticServerUrl)
	staticServerReverseProxy.Transport = &ochttp.Transport{}

	f := http.NewServeMux()
	f.HandleFunc("/", proxyHandler(staticServerReverseProxy, apiServerReverseProxy))

	log.Println("Listening proxy server on ", proxyServerPort)
	err = http.ListenAndServe(proxyServerPort, f)
	if err != nil {
		return err
	}
	done <- struct{}{}
	return nil
}

func proxyHandler(staticServerReverseProxy, apiServerReverseProxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ben", "Rad")

		if strings.Contains(r.URL.Path, "/api") {
			apiServerReverseProxy.ServeHTTP(w, r)
		} else {
			staticServerReverseProxy.ServeHTTP(w, r)
		}
	}
}
