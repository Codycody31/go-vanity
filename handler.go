// handler.go

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new router with configured routes
func NewRouter(cfg *Config) *mux.Router {
	router := mux.NewRouter()
	for _, pkg := range cfg.Packages {
		// Local copy for the closure
		path, repo := pkg.Path, pkg.Repo
		router.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprintf(w, `<meta name="go-import" content="%s git %s">`, path, repo)
		}).Methods("GET")
	}
	return router
}
