package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/hiboabd/gitline/controllers"
	"os"
)


func main() {
	port := getEnv("PORT", "1235")
	router := gin.Default()
	router.HTMLRender = controllers.CreateMyRender("web/templates")
	router.Use(static.Serve("/assets", static.LocalFile("./web/assets", true)))
	router.Use(static.Serve("/static", static.LocalFile("./web/static", true)))

	router.GET("/", controllers.RenderHomepage)
	router.GET("/timeline", controllers.GetRepositoryData)

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