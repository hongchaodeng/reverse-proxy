package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/hongchaodeng/reverse-proxy/pkg/proxy"
)

var backendAddr string
var port int

func init() {
	// parse a flag 'backend-addr' which takes a string that indicates the backend address.
	flag.StringVar(&backendAddr, "backend-addr", "", "backend server address, e.g. http://localhost:8000")
	// parse a flag 'port' which takes a number that indicates the port to listen on.
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	if backendAddr == "" {
		log.Fatal("backend-addr flag is required")
		return
	}
}

func main() {
	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := proxy.New(backendAddr)
	if err != nil {
		log.Fatalf("failed to initialize proxy: %v", err)
	}
	log.Printf("Proxy initialized with backend server: %s\n", backendAddr)

	// handle all requests to your server using the proxy
	http.HandleFunc("/", proxy.RequestHandler())
	log.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
