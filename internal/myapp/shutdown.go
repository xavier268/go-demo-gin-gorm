package myapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShutdownHdlr will trigger a gracefull server shutdown
func (a *MyApp) shutdownHdlr(c *gin.Context) {
	c.String(http.StatusOK, "Server is shutting down ...")
	fmt.Println("Server is shutting down ...")
	// Trigger shutdown asynchroneously so that the message can be returned ...
	go a.server.Shutdown(context.Background())
}
