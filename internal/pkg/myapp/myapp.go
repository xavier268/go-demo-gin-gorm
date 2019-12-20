package myapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MyApp is the main application object
// It makes the server explicitly available, to allow for graceful shutdown.
type MyApp struct {
	*gin.Engine
	server *http.Server
}

// New constructs a new MyApp application
func New() *MyApp {

	a := new(MyApp)
	a.Engine = gin.Default()
	a.server = &http.Server{
		Addr:    ":8080",
		Handler: a.Engine,
	}

	a.GET("/ping", a.pingHdlr)
	a.GET("/ping/:msg", a.pingMsgHdlr)
	a.GET("/sleep", a.pingLongHdlr)

	a.GET("/quit", a.shutdownHdlr)

	return a
}

// Run the application
func (a *MyApp) Run() {
	a.server.ListenAndServe()
}
