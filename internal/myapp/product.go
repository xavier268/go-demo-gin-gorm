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

// retrieve all products
func (a *MyApp) getAllProductsHdlr(c *gin.Context) {
	ps := a.source.AllProducts()
	c.JSON(http.StatusOK, ps)
}

// getProductCount retrieve the product count.
func (a *MyApp) getProductCountHdlr(c *gin.Context) {
	// Writing to the writer directly will trigger writing also
	// an http.StatusOK header and  a text/plain content-type header
	fmt.Fprintf(c.Writer, "%d", a.source.CountProducts())
}

func (a *MyApp) getCreateProductHdlr(c *gin.Context) {
	price, err := strconv.Atoi(c.Query("price"))
	if err != nil {
		fmt.Printf("\nInvalide price : %s\n%v\n", c.Params.ByName("price"), err)
		c.Status(http.StatusBadRequest)
		return
	}
	code := c.Query("code")
	id := a.source.CreateProduct(uint(price), code)
	fmt.Fprintf(c.Writer, "Created product %d with price = %d and code = %s", id, price, code)
}
