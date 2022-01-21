package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

type PageHandler func(w http.ResponseWriter, r *http.Request) error

const (
	ErrRenderingPage = "Error rendering page"
)
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
	wrap := errorHandler()

	router := mux.NewRouter()
	router.Handle("/",
		wrap(renderHomepage(client, templates["home.gotmpl"])))

	router.Handle("/timeline/{username}",
		wrap(
			renderTimeline(client, templates["timeline.gotmpl"])))

	static := http.FileServer(http.Dir(webDir + "/static"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", static))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", static))
	router.PathPrefix("/javascript/").Handler(http.StripPrefix("/javascript/", static))

	return router
}

func errorHandler() func(pageHandler PageHandler) http.Handler {
	return func(pageHandler PageHandler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := pageHandler(w, r)

			if err != nil {
				if redirect, ok := err.(Redirect); ok {
					http.Redirect(w, r, redirect.To(), http.StatusFound)
					return
				}

				fmt.Printf("Error handling request: %s\n", err)
				http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
				return
			}
		})
	}
}
