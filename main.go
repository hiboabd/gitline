package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := getEnv("PORT", "1235")
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Use(static.Serve("/assets", static.LocalFile("./web/assets", true)))
	router.Use(static.Serve("/static", static.LocalFile("./web/static", true)))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Gitline",
		})
	})

	//context := &gin.Context{}
	//router.GET("/", controllers.RenderHomepage(context))

	err := router.Run(":" + port)
	if err != nil {
		return
	}
}

func getEnv(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return def
}