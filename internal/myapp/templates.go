package myapp

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// initTemplates will consolidate and load the templates in the gin engine.
// Reimplement using files/glog loading as needed.
func (a *MyApp) initTemplates() {

	tpl := template.Must(template.New("DO_NOT_USE").ParseGlob("./templates/*.html"))
	fmt.Printf("\nAvailable declared templates are  : ")
	for _, t := range tpl.Templates() {
		fmt.Printf("%s, ", t.Name())
	}
	fmt.Println("")
	a.SetHTMLTemplate(tpl)
}

// This is a test handler to display template content.
func (a *MyApp) htmlHdlr(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", "hello world !")
}
