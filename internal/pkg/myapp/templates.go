package myapp

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *MyApp) initTemplates() {
	const t = `
	{{define "test1"}}
		<div>This is test1 template with message : {{.}} </div>
		{{end}}

	{{define "header"}}
		<html><body>
		{{end}}
	
	{{define "footer"}}
		</body></html>
		{{end}}
	
	{{define "main" }}
		{{template "header" .}}
		{{template "test1" . }}
		{{template "footer" . }}
		{{end}}
		
	`
	a.SetHTMLTemplate(template.Must(template.New("main").Parse(t)))
}

// This is a test handler to display template content.
func (a *MyApp) htmlHdlr(c *gin.Context) {
	c.HTML(http.StatusOK, "main", "hello world !")
}
