package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomepageHandler(t *testing.T) {

	r := getRouter()

	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/")

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	/*
		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		if err != nil {
			t.Fatal(err)
		}


		respString := string(b)

		expected := "Hello World!"


		if respString != expected {
			t.Errorf("Response should be %s, got %s", expected, respString)
		}
	*/
}
