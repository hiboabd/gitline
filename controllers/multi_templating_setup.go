package controllers

import (
	"github.com/gin-contrib/multitemplate"
)

func CreateMyRender(templatesDir string) multitemplate.Renderer {
	indexHTMLFile := templatesDir + "/index.html"
	homeHTMLFile  := templatesDir + "/home.html"
	timelineHTMLFile  := templatesDir + "/timeline.html"

	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", indexHTMLFile, homeHTMLFile)
	r.AddFromFiles("timeline", indexHTMLFile, timelineHTMLFile)
	return r
}
