// handler.go

package main

import (
	"fmt"
	"net/http"

	"go.codycody31.dev/go-vanity/config"

	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new router with configured routes
func NewRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	for _, pkg := range cfg.Packages {
		// Local copy for the closure
		path, repo := pkg.Path, pkg.Repo
		router.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")

			// If go-get is not provided, redirect to the pgk.go.dev page
			if r.URL.Query().Get("go-get") != "1" {
				http.Redirect(w, r, "https://pkg.go.dev/"+cfg.Domain+"/"+path, http.StatusFound)
				return
			}

			fmt.Fprintf(w, `<meta name="go-import" content="%s git %s">`, cfg.Domain+"/"+path, repo)
		}).Methods("GET")
	}

	// Then, if index is requested, show the list of packages with go-get meta tags
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<html><body><h1>Packages</h1><ul>")
		for _, pkg := range cfg.Packages {
			fmt.Fprintf(w, `<li><a href="/%s">%s</a></li>`, pkg.Path, pkg.Path)
		}
		fmt.Fprintf(w, "</ul></body></html>")
	}).Methods("GET")

	return router
}
