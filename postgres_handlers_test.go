package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestPostgresConnectHandler(t *testing.T) {

	// Set up some test data - two ways of doing this
	d := Database{}
	fmt.Println(d)

	// Get a router instance
	r := getRouter()

	// Set up a mock http server and pass it our router
	mockServer := httptest.NewServer(r)

	// Call /birds URL from mockServer
	resp, err := http.Get(mockServer.URL + "/connect")
	if err != nil {
		t.Fatalf("Error is: %s", err)
	}

	// Check the status code of the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Check the Content-Type header
	expected := "text/plain; charset=utf-8"
	actual := resp.Header.Get("Content-Type")
	if expected != actual {
		t.Errorf("Problem with Content-Type header, expected %s, got %s", expected, actual)
	}

	// Let's post data and verify that this updates the Database struct accordingly
	form := connectForm()
	requestBody := form.Encode()
	postResp, _ := http.Post(mockServer.URL+"/connect", "application/x-www-form-urlencoded", bytes.NewBufferString(requestBody))

	fmt.Println(postResp)

}

func connectForm() *url.Values {
	form := url.Values{}
	form.Set("host", "jupiter")
	form.Set("port", "1234")
	form.Set("user", "postgres")
	form.Set("port", "p0stgres")
	return &form
}
