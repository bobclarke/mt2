package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

// Database struct represents our database
type Database struct {
	DBHost       string
	DBPort       string
	DBLogin      string
	DBPassword   string
	DBName       string
	DBConnection []string
	DBTables     []string
}

var d = &Database{}

func connect(w http.ResponseWriter, r *http.Request) {

	//fmt.Printf("Method is: %s \n", r.Method)
	r.ParseForm()
	d.DBHost = r.Form.Get("host")
	d.DBPort = r.Form.Get("port")
	d.DBLogin = r.Form.Get("user")
	d.DBPassword = r.Form.Get("password")

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func read(w http.ResponseWriter, r *http.Request) {
	dBytes, _ := json.Marshal(d)
	w.Write(dBytes)
}

func (d Database) getMovies() []string {

	movies := []string{"A Good Year", "The Matrix", "The Hangover"}
	fmt.Printf("Movies: %v", movies)
	return movies
}
