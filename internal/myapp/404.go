package myapp

import (
	"github.com/gin-gonic/gin"
)

func (a *MyApp) notFoundHdlr(c *gin.Context) {
	data := gin.H{
		"url": c.Request.URL,
	}
	a.render(c, data, "404.html")
}
