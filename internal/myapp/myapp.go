package myapp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/dao"
)

// MyApp is the main application object.
// It makes the server explicitly available, to allow for graceful shutdown.
type MyApp struct {
	*gin.Engine
	server *http.Server
	source *dao.Source
}

// New constructs a new MyApp application, using the provided data source.
// The source should not be nil.
func New(source *dao.Source) *MyApp {

	a := new(MyApp)
	a.Engine = gin.Default()
	a.server = &http.Server{
		Addr:    ":8080",
		Handler: a.Engine,
	}

	// Select data source
	if source == nil {
		panic("You provided a nil source to initialize the app ?!")
	}
	a.source = source

	// Initialize templates
	a.initTemplates()

	// Define routes
	a.NoRoute(a.notFoundHdlr)

	v1 := a.Group("/v1")
	{
		v1.GET("/ping", a.pingHdlr)
		v1.GET("/ping/:msg", a.pingMsgHdlr)
		v1.GET("/sleep", a.pingLongHdlr)
		v1.GET("/temp", a.htmlHdlr)

		v1.GET("/p/:id", a.getProductHdlr)
		v1.GET("/ps/", a.getAllProductsHdlr)
		v1.GET("/count", a.getProductCountHdlr)

		// price and code expected in query ...
		v1.GET("/create", a.getCreateProductHdlr)

		v1.GET("/quit", a.shutdownHdlr)
	}

	// Dump routes
	a.DumpRoutes()

	return a
}

// Run the application, blocking call.
func (a *MyApp) Run() {
	fmt.Printf(
		"\n========================\n"+
			"Serving from address  %s"+
			"\n========================\n",
		a.server.Addr)
	err := a.server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// Shutdown application, closing the data source.
func (a *MyApp) Shutdown() {
	go a.server.Shutdown(context.Background())
	if a.source != nil {
		a.source.Close()
	}
}

// DumpRoutes dumps the registrered routes.
func (a *MyApp) DumpRoutes() {
	fmt.Printf("\n========\nRoutes\n========\n")
	for _, r := range a.Routes() {
		fmt.Printf("%s\t%s\n\t===>\t%s\n", r.Method, r.Path, r.Handler)
	}
	fmt.Printf("========\n")
}
