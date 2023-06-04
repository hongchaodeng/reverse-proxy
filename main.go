package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hongchaodeng/reverse-proxy/pkg/proxy"
)

var routeStr string
var port int

func init() {
	flag.StringVar(&routeStr, "routes", "/:http://localhost:8080", "comma separated routing rules, each of format $path:$url, e.g. /example:http://api-server:8000")
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	if routeStr == "" {
		log.Fatal("backend-addr flag is required")
		return
	}
}

func main() {
	// initialize a reverse proxy and pass the actual backend server url here
	p, err := proxy.New(routeStr)
	if err != nil {
		log.Fatalf("failed to initialize proxy: %v", err)
	}

	addr := fmt.Sprintf(":%d", port)
	log.Fatal(p.Start(addr))
}
