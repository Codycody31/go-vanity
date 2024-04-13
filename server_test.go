// server_test.go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	cfg := &Config{
		Packages: []Package{
			{Path: "go.example.com/mylib", Repo: "https://github.com/username/mylib"},
		},
	}

	router := NewRouter(cfg)
	server := httptest.NewServer(router)
	defer server.Close()

	// Test if the server responds with the correct meta tag
	res, err := http.Get(server.URL + "/go.example.com/mylib")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}
}
