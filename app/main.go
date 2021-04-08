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
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	fmt.Println("Starting server...")

	development := true
	if strings.EqualFold(os.Getenv("GIN_MODE"), "release") {
		development = false
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	var db *gorm.DB
	if strings.EqualFold(os.Getenv("DBMS"), "mysql") {
		db = database.MysqlConnect(development)
	} else {
		db = database.PostgresConnect(development)
	}

	middleware.InitErrorHandler(r)
	module.Init(db, r, development)

	//Development Endpoint
	if development {
		module.Development(db, r)
	}
	r.Run()

}
