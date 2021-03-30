package main

import (
	"fmt"

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
	r := gin.Default()
	middleware.InitErrorHandler(r)
	db := database.PostgresConnect()
	module.Init(db, r)
	r.Run()

}
