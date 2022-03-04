package main

import (
	"context"
	"fmt"
	"github.com/hiboabd/gitline/internal/controllers"
	"github.com/hiboabd/gitline/internal/handlers"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	port := handlers.GetEnv("PORT", "1235")
	webDir := handlers.GetEnv("WEB_DIR", "web")
	apiUrl := handlers.GetEnv("API_URL", "")

	fmt.Println("port", port)
	fmt.Println("web dir", webDir)
	fmt.Println("api url", apiUrl)

	layouts, _ := template.
		New("").
		Funcs(map[string]interface{}{}).
		ParseGlob(webDir + "/templates/*.gotmpl")

	files, _ := filepath.Glob(webDir + "/templates/*.gotmpl")
	tmpls := map[string]*template.Template{}

	for _, file := range files {
		tmpls[filepath.Base(file)] = template.Must(template.Must(layouts.Clone()).ParseFiles(file))
	}

	client, err := handlers.NewClient(http.DefaultClient, apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	s := &http.Server{
		Addr:    ":" + port,
		Handler: controllers.New(client, tmpls, webDir),
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("Running at :" + port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	sig := <-c
	fmt.Println("signal received: ", sig)


	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.Shutdown(tc); err != nil {
		fmt.Println(err)
	}
}