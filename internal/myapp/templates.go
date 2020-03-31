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

	tt := []string{
		"./templates/*.html",
		"../templates/*.html",
		"../../templates/*.html",
		"../../../templates/*.html",
	}

	for _, tempPath := range tt {
		tpl, err := template.New("DO_NOT_USE").ParseGlob(tempPath)
		if err == nil {
			fmt.Println("Loaded templates from : ", tempPath)
			fmt.Printf("\nAvailable declared templates are  : ")
			for _, t := range tpl.Templates() {
				fmt.Printf("%s, ", t.Name())
			}
			fmt.Println("")
			a.SetHTMLTemplate(tpl)
			return
		}
		fmt.Println("Tried to load templates from : ", tempPath, "\n-->\terror :", err)
	}
	panic("Could not load templates ")

}

// This is a test handler to display template content.
func (a *MyApp) htmlHdlr(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", "hello world !")
}
