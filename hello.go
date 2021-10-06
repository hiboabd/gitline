package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func main() {
	// creates a router object with built-in defaults that come with gin
	router := gin.Default()

	// assign handler function to be called for any http get requests to /hello
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// create a group of routes behind the path /api
	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// enables webserver to serve any static files from the views directory
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// start web server
	err := router.Run()
	if err != nil {
		return
	}
}