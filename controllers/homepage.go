package controllers

import (
	"net/http"
)

func RenderHomepage(w http.ResponseWriter, r *http.Request) (string, interface{}, error) {
	return "home.html", nil, nil
}