package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/paradewisudaitb/Backend/connection/database"
	"github.com/paradewisudaitb/Backend/module"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting server...")

	development := true
	if strings.EqualFold(os.Getenv("GIN_MODE"), "release") {
		development = false
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	db := database.PostgresConnect(development)

	middleware.InitErrorHandler(r)
	module.Init(db, r)

	//Development Endpoint
	module.Development(db, r)
	r.Run()

}
