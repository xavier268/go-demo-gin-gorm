package myapp

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/dao"
)

// retrieve product from its id
func (a *MyApp) getProductHdlr(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println("Error retieving product : ", err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	p := dao.GetDAO().GetProduct(uint(id))
	c.JSON(http.StatusOK, *p)
}
