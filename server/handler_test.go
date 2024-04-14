// server_test.go

package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.codycody31.dev/vanity/config"
)

func TestServer(t *testing.T) {
	cfg := &config.Config{
		Domain: "go.example.com",
		Packages: []config.Package{
			{Path: "mylib", Repo: "https://github.com/username/mylib"},
		},
	}

	router := NewRouter(cfg)
	server := httptest.NewServer(router)
	defer server.Close()

	// Test if the server responds with the correct meta tag
	res, err := http.Get(server.URL + "/mylib")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", res.StatusCode)
	}
}
