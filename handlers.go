package main

import (
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func GenerateShortUrlHandler(w http.ResponseWriter, r *http.Request) {
	// TODO persistence
	// TODO shortUrl Lifetime ?
	// TODO Analytics (threading)
	url := mux.Vars(r)["url"]
	hasher := new(Hasher)
	hStr := hasher.GenerateHashFromURL(url)
	siteHashes[hStr] = url
	result := r.Host +"/" + hStr
	io.WriteString(w, result)
}

func ForwardHandler(w http.ResponseWriter, r *http.Request) {
	// TODO handle non-present hash
	path := mux.Vars(r)["hash"]
	w.Header().Set("Location",siteHashes[path])
	w.WriteHeader(http.StatusMovedPermanently)
}