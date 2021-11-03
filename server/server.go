package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/repositories", GetRepos).Methods("GET")

	fmt.Println("Mock server running at port 2000")
	log.Fatal(http.ListenAndServe(":2000", r))
}
