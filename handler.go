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

			// If go-get is not provided, redirect to the pgk.go.dev page
			if r.URL.Query().Get("go-get") != "1" {
				http.Redirect(w, r, "https://pkg.go.dev/"+cfg.Domain+"/"+path, http.StatusFound)
				return
			}

			fmt.Fprintf(w, `<meta name="go-import" content="%s git %s">`, cfg.Domain+"/"+path, repo)
		}).Methods("GET")
	}
	return router
}
