package config

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromFile(t *testing.T) {
	// Create a temporary YAML file
	tempFile, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		t.Fatalf("Cannot create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // clean up

	content := []byte("domain: go.example.com\npackages:\n  - path: mylib\n    repo: https://github.com/username/mylib\n    vcs: git")
	if _, err := tempFile.Write(content); err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Failed to close file: %v", err)
	}

	config, err := LoadConfig(tempFile.Name(), "")
	assert.NoError(t, err)
	assert.Equal(t, "go.example.com", config.Domain)
	assert.Len(t, config.Packages, 1)
	assert.Equal(t, "mylib", config.Packages[0].Path)
}

func TestLoadConfigFromURL(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("domain: go.example.com\npackages:\n  - path: mylib\n    repo: https://github.com/username/mylib\n    vcs: git"))
		if err != nil {
			t.Fatalf("Failed to write response: %v", err)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	config, err := LoadConfig("", server.URL)
	assert.NoError(t, err)
	assert.Equal(t, "go.example.com", config.Domain)
	assert.Len(t, config.Packages, 1)
	assert.Equal(t, "mylib", config.Packages[0].Path)
}
