package module

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORSModule(g *gin.Engine, devmode bool) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://wisprilitb.com", "https://staging.wisprilitb.com"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}
	g.Use(cors.New(corsConfig))
}
