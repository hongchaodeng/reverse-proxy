package proxy

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// Proxy is an interface that provides a new reverse proxy.
type Proxy interface {
	// Start starts a proxy server that listens on given addr.
	// Start blocks until the server is stopped or an error occurs.
	// Start always returns a non-nil error.
	Start(string) error
}

type proxyimpl struct {
	pathToTarget map[string]string
}

func (p *proxyimpl) Start(addr string) error {
	for path, target := range p.pathToTarget {
		log.Printf("routing from %s to %s", path, target)
		u, err := url.Parse(target)
		if err != nil {
			return err
		}
		// HandleFunc takes a copy of the path because it is async
		pathVar := path
		http.HandleFunc(pathVar, func(w http.ResponseWriter, r *http.Request) {
			// Update the headers to allow for SSL redirection
			r.Host = u.Host
			r.URL.Host = u.Host
			r.URL.Scheme = u.Scheme

			// trim path prefix
			urlpath := r.URL.Path
			r.URL.Path = strings.TrimLeft(urlpath, pathVar)

			proxy := httputil.NewSingleHostReverseProxy(u)
			proxy.ServeHTTP(w, r)
		})
	}
	log.Printf("Starting server on addr %s\n", addr)
	return http.ListenAndServe(addr, nil)
}

// NewProxy takes routing rules and creates a reverse proxy
func New(routeStr string) (Proxy, error) {
	pathToTarget := make(map[string]string)
	for _, route := range strings.Split(routeStr, ",") {
		// split 'route' based on the first ':' character
		// the first part is the path, the second part is the url
		pairs := strings.SplitN(route, ":", 2)
		if len(pairs) != 2 {
			return nil, fmt.Errorf("invalid route: %s", route)
		}
		pathToTarget[pairs[0]] = pairs[1]
	}
	return &proxyimpl{
		pathToTarget: pathToTarget,
	}, nil
}
