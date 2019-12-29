package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/dao"
	"github.com/xavier268/go-demo-gin-gorm/internal/myapp"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	fmt.Println("Type Ctrl-C to stop the server")

	src := dao.NewPostgresSource()
	defer src.Close()

	myapp.New(src).Run()
}
