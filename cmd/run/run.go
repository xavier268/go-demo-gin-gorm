package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xavier268/go-demo-gin-gorm/internal/pkg/myapp"
)

func main() {

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	fmt.Println("Type Ctrl-C to stop the server")

	myapp.New().Run()
}
