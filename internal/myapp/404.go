package myapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *MyApp) notFoundHdlr(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", c.Request.URL)
}
