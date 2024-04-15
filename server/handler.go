// handler.go

package server

import (
	"fmt"
	"net/http"

	"go.codycody31.dev/vanity/config"

	"github.com/gorilla/mux"
)

// NewRouter creates and returns a new router with configured routes
func NewRouter(cfg *config.Config) *mux.Router {
	router := mux.NewRouter()
	for _, pkg := range cfg.Packages {
		// Local copy for the closure
		path, repo, vcs := pkg.Path, pkg.Repo, pkg.VCS
		router.HandleFunc("/"+path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")

			// If go-get is not provided, redirect to the pgk.go.dev page
			if r.URL.Query().Get("go-get") != "1" {
				http.Redirect(w, r, "https://pkg.go.dev/"+cfg.Domain+"/"+path, http.StatusFound)
				return
			}

			fmt.Fprintf(w, `<meta name="go-import" content="%s %s %s">`, cfg.Domain+"/"+path, vcs, repo)
		}).Methods("GET")
	}

	// Then, if index is requested, show the list of packages with go-get meta tags
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<title>Packages</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			margin: 40px;
			background-color: #f4f4f9;
			color: #333;
		}
		h1 {
			color: #5a5a5a;
		}
		ul {
			list-style-type: none;
			padding: 0;
		}
		li {
			margin: 10px 0;
			padding: 10px;
			background-color: #fff;
			border: 1px solid #ddd;
		}
		a {
			text-decoration: none;
			color: #0066cc;
		}
		a:hover {
			text-decoration: underline;
		}
	</style>
</head>
<body>
	<h1>Available Packages</h1>
	<ul>
	`)
		for _, pkg := range cfg.Packages {
			fmt.Fprintf(w, `<li><a href="/%s">%s</a></li>`, pkg.Path, pkg.Path)
		}
		fmt.Fprintf(w, "</ul></body></html>")
	}).Methods("GET")

	return router
}
