package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	port := flag.String("port", "8080", "Port for the proxy to listen on")
	backend := flag.String("backend", "http://localhost:11434", "Backend URL to proxy to")
	flag.Parse()

	target, err := url.Parse(*backend)
	if err != nil {
		log.Fatalf("invalid backend URL: %v", err)
	}

	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
			req.RequestURI = ""

			// Remove Origin (CORS) header as Ollama returns 403 otherwise.
			req.Header.Del("Origin")
		},
		ModifyResponse: func(resp *http.Response) error {
			return nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			log.Printf("proxy error: %v", err)
			http.Error(w, "Proxy error", http.StatusBadGateway)
		},
	}

	addr := ":" + *port
	log.Printf("Proxy listening on http://localhost:%s, forwarding to %s", *port, target.String())
	if err := http.ListenAndServe(addr, proxy); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
