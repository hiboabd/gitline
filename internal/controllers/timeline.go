package controllers

import (
	"github.com/gorilla/mux"
	"github.com/hiboabd/gitline/internal/handlers"
	"net/http"
)

type UsernameInformation interface {
	GetRepositoryData(username string) (handlers.RepositoryData, error)
}

type listReposVars struct {
	Path         string
	XSRFToken    string
	UserRepositories handlers.RepositoryData
	Error        string
	ErrorMessage string
}

func renderTimeline(client UsernameInformation, tmpl Template) PageHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != http.MethodGet {
			return StatusError(http.StatusMethodNotAllowed)
		}

		routeVars := mux.Vars(r)
		username := routeVars["username"]
		userRepositories, err := client.GetRepositoryData(username)
		if err != nil {
			return err
		}

		vars := listReposVars{
			Path:         r.URL.Path,
			UserRepositories: userRepositories,
		}

		return tmpl.ExecuteTemplate(w, "index", vars)
	}
}
