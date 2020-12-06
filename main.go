package main

import (
	"github.com/gorilla/mux"
	"hash/fnv"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

var siteHashes = make(map[string]string)

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/{hash}", forward).Methods("GET").Name("router")
	r.HandleFunc("/", generateShortUrlHandler).Methods("POST").Queries("url", "{url}")
	return r
}

func forward(w http.ResponseWriter, r *http.Request) {
	// TODO handle non-present hash
	path := mux.Vars(r)["hash"]
	w.Header().Set("Location",siteHashes[path])
	w.WriteHeader(http.StatusMovedPermanently)
}

var strconvFormatUint = strconv.FormatUint

func generateShortUrl(url string) (string) {
	h := hash(url)
	var hStr = strconvFormatUint(uint64(h), 10)
	return hStr
}

func generateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	// TODO persistence
	// TODO shortUrl Lifetime ?
	// TODO Analytics (threading)
	url := mux.Vars(r)["url"]
	hStr := generateShortUrl(url)
	siteHashes[hStr] = url
	result := r.Host +"/" + hStr
	io.WriteString(w, result)
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

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
