package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

type PageHandler func(w http.ResponseWriter, r *http.Request) (templateName string, data interface{}, err error)

const (
	ErrRenderingPage = "Error rendering page"
)

var templates = map[string]*template.Template{}

func Register() {
	layouts, _ := template.
		New("").
		Funcs(map[string]interface{}{}).
		ParseGlob("/web/templates/*.html")

	files, _ := filepath.Glob("/web/templates/*.html")

	for _, file := range files {
		templates[filepath.Base(file)] = template.Must(template.Must(layouts.Clone()).ParseFiles(file))
	}
}

func Render(h PageHandler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		templateName, data, err := h(w, r)
		if err != nil {
			fmt.Printf("Error handling request: %s\n", err)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}

		tmpl, ok := templates[templateName]
		if !ok {
			fmt.Printf("No tmpl not found: %s\n", templateName)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}

		var b bytes.Buffer
		err = tmpl.Execute(&b, data)
		if err != nil {
			fmt.Printf("Error executing template: %s\n", err)
			http.Error(w, ErrRenderingPage, http.StatusInternalServerError)
			return
		}
		io.Copy(w, &b)
	}
	return fn
}
