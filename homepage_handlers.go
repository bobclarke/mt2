package main

import (
	"net/http"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
