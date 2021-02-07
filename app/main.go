package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func main() {
	fmt.Println("Starting server...")
	r := gin.Default()
	r.Run()

}
