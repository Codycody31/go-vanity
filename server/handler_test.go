package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.codycody31.dev/vanity/config"

	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	cfg := &config.Config{
		Domain: "go.example.com",
		Packages: []config.Package{
			{Path: "mypackage", Repo: "https://github.com/example/mypackage", VCS: "git"},
			{Path: "myotherpackage", Repo: "https://github.com/example/myotherpackage", VCS: "git"},
		},
	}

	router := NewRouter(cfg)

	tests := []struct {
		route   string
		method  string
		query   string
		status  int
		content string
	}{
		{"/mypackage", "GET", "go-get=1", http.StatusOK, `<meta name="go-import" content="go.example.com/mypackage git https://github.com/example/mypackage">`},
		{"/mypackage", "GET", "", http.StatusFound, ""},
		{"/", "GET", "", http.StatusOK, "<li><a href=\"/mypackage\">mypackage</a> (<a href=\"https://github.com/example/mypackage\" class="},
		{"/nonexistent", "GET", "", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.route, nil)
		if tt.query != "" {
			req.URL.RawQuery = tt.query
		}
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		assert.Equal(t, tt.status, rr.Code, fmt.Sprintf("Route %s failed", tt.route))
		if tt.content != "" {
			assert.Contains(t, rr.Body.String(), tt.content, fmt.Sprintf("Route %s returned unexpected body", tt.route))
		}
	}
}
