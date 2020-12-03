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
	r.HandleFunc("/", generateShortUrl).Methods("POST").Queries("url", "{url}")
	return r
}

func forward(w http.ResponseWriter, r *http.Request) {
	// TODO handle non-present hash
	path := mux.Vars(r)["hash"]
	w.Header().Set("Location",siteHashes[path])
	w.WriteHeader(http.StatusMovedPermanently)
}

func generateShortUrl(w http.ResponseWriter, r *http.Request) {
	// TODO persistence
	// TODO shortUrl Lifetime ?
	// TODO Analytics (threading)
	url := mux.Vars(r)["url"]
	h := hash(url)
	var hStr = strconv.FormatUint(uint64(h), 10)
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
