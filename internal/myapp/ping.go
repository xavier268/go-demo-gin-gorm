package myapp

import (
	"time"

	"github.com/gin-gonic/gin"
)

// pingHdlr emits a pong
func (a *MyApp) pingHdlr(c *gin.Context) {
	data := gin.H{
		"Ping": "pong",
	}
	a.render(c, data, "ping.html")

}

// pingMsgHdlr emits the receives message from the url
func (a *MyApp) pingMsgHdlr(c *gin.Context) {
	data := gin.H{
		"Ping": c.Params.ByName("msg"),
	}
	a.render(c, data, "ping.html")
}

// pingLongHdlr is a LOOOONG ping request ...
func (a *MyApp) pingLongHdlr(c *gin.Context) {
	time.Sleep(1 * time.Second)
	data := gin.H{
		"Ping": "sleep 1 second",
	}
	a.render(c, data, "ping.html")
}
