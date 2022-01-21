package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{username}/repos", GetRepos).Methods("GET")

	fmt.Println("Mock server running at port 2000")
	log.Fatal(http.ListenAndServe(":2000", r))
}
