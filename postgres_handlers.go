package main

import (
	"encoding/json"
	"net/http"
)

// Database struct represents our database
type Database struct {
	Host      string
	Port      int64
	User      string
	Password  string
	Connected bool
	Databases []string
}

func connect(w http.ResponseWriter, r *http.Request) {

	// Get a pointer to the Database struct
	d := &Database{}

	// Update the connected status of our database
	d.Connected = true

	// Generate a JSON version of our Database struct
	dBytes, _ := json.Marshal(d)

	// Write the json to via our ResponseWriter
	w.Write(dBytes)
}
