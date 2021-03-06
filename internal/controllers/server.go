package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

type PageHandler func(w http.ResponseWriter, r *http.Request) error

type Redirect string

func (e Redirect) Error() string {
	return "redirect to " + string(e)
}

func (e Redirect) To() string {
	return string(e)
}

type StatusError int

func (e StatusError) Error() string {
	code := e.Code()

	return fmt.Sprintf("%d %s", code, http.StatusText(code))
}

func (e StatusError) Code() int {
	return int(e)
}

type Client interface {
	UsernameInformation
}

type Template interface {
	ExecuteTemplate(io.Writer, string, interface{}) error
}

func New(client Client, templates map[string]*template.Template, webDir string) http.Handler {
	wrap := errorHandler(templates["error.gotmpl"])

	router := mux.NewRouter()
	router.Handle("/",
		wrap(renderHomepage(client, templates["home.gotmpl"])))

	router.Handle("/timeline/{username}",
		wrap(
			renderTimeline(client, templates["timeline.gotmpl"])))

	static := http.FileServer(http.Dir(webDir + "/static"))
	router.PathPrefix("/assets/").Handler(static)
	router.PathPrefix("/javascript/").Handler(static)
	router.PathPrefix("/stylesheets/").Handler(static)
	router.PathPrefix("/images/").Handler(static)

	return http.StripPrefix("", router)
}

type errorVars struct {
	Code      int
	Error     string
}

func errorHandler(tmplError Template) func(pageHandler PageHandler) http.Handler {
	return func(pageHandler PageHandler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := pageHandler(w, r)

			if err != nil {
				if redirect, ok := err.(Redirect); ok {
					http.Redirect(w, r, redirect.To(), http.StatusFound)
					return
				}

				code := http.StatusInternalServerError
				if status, ok := err.(StatusError); ok {
					if status.Code() == http.StatusForbidden || status.Code() == http.StatusNotFound {
						code = status.Code()
					}
				}

				w.WriteHeader(code)
				err = tmplError.ExecuteTemplate(w, "index", errorVars{
					Code:      code,
					Error:     err.Error(),
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
			}
		})
	}
}
