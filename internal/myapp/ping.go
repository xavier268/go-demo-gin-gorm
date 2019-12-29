package myapp

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/models"
)

// pingHdlr emits a pong
func (a *MyApp) pingHdlr(c *gin.Context) {
	ping := new(models.Ping)
	ping.Ping = "Pong"
	c.JSON(200, ping)
}

// pingMsgHdlr emits the receives message from the url
func (a *MyApp) pingMsgHdlr(c *gin.Context) {
	m := new(models.Ping)
	m.Ping = c.Params.ByName("msg")
	c.JSON(200, m)
}

// pingLongHdlr is a LOOOONG ping request ...
func (a *MyApp) pingLongHdlr(c *gin.Context) {
	time.Sleep(5 * time.Second)
	m := new(models.Ping)
	m.Ping = "Sleep"
	c.JSON(200, m)
}
