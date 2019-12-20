package myapp

import "github.com/gin-gonic/gin"

import "net/http"

func (a *MyApp) notFoundHdlr(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", c.Request.URL)
}
