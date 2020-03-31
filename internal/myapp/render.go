package myapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// render will render the data as json, xml, or htm (with provided template).
// selection is done based upon the Accept header, if present, or the existence
// of a query parameter named json or xml.
func (a *MyApp) render(c *gin.Context, data gin.H, templateName string) {
	h := c.Request.Header.Get("Accept")
	switch {
	case h == "application/json" || c.Query("json") != "":
		{
			// respond with json
			c.IndentedJSON(http.StatusOK, data)
		}
	case h == "application/xml" || c.Query("xml") != "":
		{
			// respond with xml
			c.XML(http.StatusOK, data)
		}
	case templateName != "":
		{
			// respond with html
			c.HTML(http.StatusOK, templateName, data)
		}
	default:
		{
			// send string error message
			c.String(http.StatusNotAcceptable, "html rendering is not acceptable for this path")
		}
	}
}
