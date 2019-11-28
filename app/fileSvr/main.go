package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rs/cors"
)

var d = flag.String("d", "./", "the directory")
var host = flag.String("host", ":8080", "the host to be bind")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*d)))

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	handler := cors.AllowAll().Handler(mux)
	log.Println("file server listen on:", *host)
	http.ListenAndServe(*host, handler)
}
