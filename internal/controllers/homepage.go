package controllers

import (
	"fmt"
	"net/http"
)

func renderHomepage(_ Client, tmpl Template) PageHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		switch r.Method {
		case http.MethodGet:
			return tmpl.ExecuteTemplate(w, "index", nil)
		case http.MethodPost:
			username := r.PostFormValue("username")
			return Redirect(fmt.Sprintf("/timeline/%s", username))
		default:
			return StatusError(http.StatusMethodNotAllowed)
		}
	}
}