package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

var siteHashes = make(map[string]string)

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{hash}", ForwardHandler).Methods("GET").Name("router")
	r.HandleFunc("/", GenerateShortUrlHandler).Methods("POST").Queries("url", "{url}")
	return r
}


// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
