package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

var assetsDir = "/assets/"

func main() {
	// Get a router instance and start our HTTP server
	r := getRouter()
	http.ListenAndServe("127.0.0.1:8000", r)
}

func getRouter() *mux.Router {

	r := mux.NewRouter()

	// Assign handlers
	r.HandleFunc("/connect", connect).Methods("GET")
	r.HandleFunc("/", homepage).Methods("GET")

	// Static content handler (i.e for stuff in /assets/)
	staticContent := http.Dir("." + assetsDir)
	staticFileHandler := http.StripPrefix(assetsDir, http.FileServer(staticContent))
	r.PathPrefix(assetsDir).Handler(staticFileHandler).Methods("GET")

	return r
}
