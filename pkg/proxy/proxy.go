package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

// Proxy is an interface that provides a new reverse proxy.
//
// Example:
//
//	p, err := New("http://localhost:8080")
//	if err != nil {
//	  log.Fatal(err)
//	}
//	http.HandleFunc("/", p.RequestHandler())
//	http.ListenAndServe(":9090", nil)
type Proxy interface {
	// RequestHandler returns a function that can be used as an http.Handler.
	RequestHandler() func(http.ResponseWriter, *http.Request)
}

type proxyimpl struct {
	rp *httputil.ReverseProxy
}

func (p *proxyimpl) RequestHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		p.rp.ServeHTTP(w, r)
	}
}

// NewProxy takes target host and creates a reverse proxy
func New(target string) (Proxy, error) {
	url, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(url)

	return &proxyimpl{
		rp: rp,
	}, nil
}
