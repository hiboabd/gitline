package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hiboabd/gitline/controllers"
	"github.com/hiboabd/gitline/render"
	"html/template"
	"net/http"
	"os"
)

func Configure() http.Handler {
	render.Register("home.html",
		template.Must(template.ParseFiles( "web/templates/index.html", "web/templates/home.html")),
	)
	render.Register("timeline.html",
		template.Must(template.ParseFiles("web/templates/index.html", "web/templates/timeline.html")),
	)

	apiUrl := getEnv("API_URL", "")
	client, err := controllers.NewClient(http.DefaultClient, apiUrl)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
	router.HandleFunc("/", render.Render(client.RenderHomepage))
	router.HandleFunc("/timeline", render.Render(client.GetRepositoryData))

	return router
}

func main() {
	port := getEnv("PORT", "1235")
	err := http.ListenAndServe(":" + port, Configure())
	if err != nil {
		panic(err)
	}
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}