package myapp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// retrieve product from its id
func (a *MyApp) getProductHdlr(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println("Error retrieving product id : ", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	p := a.source.GetProduct(uint(id))
	c.JSON(http.StatusOK, *p)
}
