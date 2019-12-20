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
	const t = `
	{{define "test"}}
		<div>This is test template with message : {{.}} </div>
		{{end}}

	{{define "header"}}
		<html><body>
		{{end}}
	
	{{define "footer"}}
		</body></html>
		{{end}}
	
	{{define "main" }}
		{{template "header" .}}
		{{template "test" . }}
		{{template "footer" . }}
		{{end}}
		
	`
	tpl := template.Must(template.New("DO_NOT_USE").Parse(t))
	fmt.Printf("\nAvailable declared templates are  : ")
	for _, t := range tpl.Templates() {
		fmt.Printf("%s, ", t.Name())
	}
	fmt.Println("")
	a.SetHTMLTemplate(tpl)
}

// This is a test handler to display template content.
func (a *MyApp) htmlHdlr(c *gin.Context) {
	c.HTML(http.StatusOK, "main", "hello world !")
}
