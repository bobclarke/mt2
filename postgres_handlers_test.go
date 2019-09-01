package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConnect(t *testing.T) {

	// Set up some test data - two ways of doing this
	d := Database{}
	d.connected = true

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

	// Check the response body make make sure it contains our test data
	/*
		expectedRespBody, _ := json.Marshal(testBird) // Convert from type Bird to JSON
		actualRespBody, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		x := actualRespBody[:len(actualRespBody)-1]
		trimmedRespBody := x[1:]

		if string(expectedRespBody) != string(trimmedRespBody) {
			t.Errorf("Expected response of: %s, got %s", expectedRespBody, trimmedRespBody)
		}
	*/
}
