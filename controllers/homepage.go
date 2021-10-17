package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RenderHomepage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"index.tmpl",
		gin.H{
			"title": "Gitline",
		},
	)
}