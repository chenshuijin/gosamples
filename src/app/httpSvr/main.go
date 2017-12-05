package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	defaultSample()
}

func muxSample() {
	r := mux.NewRouter()
	r.Headers("Content-Type", "application/(text|json)")
	r.HandleFunc("/", DefaultRoute).Methods("GET")
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	log.Println("muxSample listen on 8000")
	// Bind to a port and pass router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

func defaultSample() {
	log.Println("defaultSample listen on 8000")
	http.HandleFunc("/", DefaultRoute)
	// Bind to a port and pass router in
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome here!"))
}
