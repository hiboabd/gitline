package main

import (
	"github.com/gorilla/mux"
	"github.com/hiboabd/gitline/controllers"
	"github.com/hiboabd/gitline/render"
	"net/http"
	"os"
)

func Configure() http.Handler {
	render.Register()
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets"))))
	router.HandleFunc("/", render.Render(controllers.RenderHomepage))
	router.HandleFunc("/timeline", render.Render(controllers.GetRepositoryData))

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