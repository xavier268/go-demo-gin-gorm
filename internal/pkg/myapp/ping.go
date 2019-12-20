package myapp

import "github.com/gin-gonic/gin"

import "time"

// pingHdlr emits a pong
func (a *MyApp) pingHdlr(c *gin.Context) {
	c.JSON(200, gin.H{
		"ping": "pong",
	})
}

// pingMsgHdlr emits the receives message from the url
func (a *MyApp) pingMsgHdlr(c *gin.Context) {
	m := c.Params.ByName("msg")
	c.JSON(200, gin.H{
		"message": m,
	})
}

// pingLongHdlr is a LOOOONG ping request ...
func (a *MyApp) pingLongHdlr(c *gin.Context) {
	time.Sleep(30 * time.Second)
	c.JSON(200, gin.H{
		"delayed": "30 seconds",
	})
}
