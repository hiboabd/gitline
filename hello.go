package main

import "github.com/gin-gonic/gin"

func main() {
	// creates a router object with built-in defaults that come with gin
	router := gin.Default()

	// assign handler function to be called for any http get requests to /hello
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// start web server on port 3000
	err := router.Run(":3001")
	if err != nil {
		return
	}
}