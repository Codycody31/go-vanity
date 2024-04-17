// handler.go

package server

import (
	"fmt"
	"net/http"

	"go.codycody31.dev/vanity/config"
	"go.codycody31.dev/vanity/templates"

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

			// If go-get is not provided, redirect to the pkg.go.dev page
			if r.URL.Query().Get("go-get") != "1" {
				http.Redirect(w, r, "https://pkg.go.dev/"+cfg.Domain+"/"+path, http.StatusFound)
				return
			}

			// Response for go-get query
			_, err := fmt.Fprintf(w, `<meta name="go-import" content="%s %s %s">`, cfg.Domain+"/"+path, vcs, repo)
			if err != nil {
				return
			}
		}).Methods("GET")
	}

	if !cfg.DisableRootPackagesPage {
		// If index is requested, show the list of packages with links to both package details and repository
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			err := templates.Templates.ExecuteTemplate(w, "index.html", map[string]interface{}{
				"Packages": cfg.Packages,
			})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})
	}

	return router
}
