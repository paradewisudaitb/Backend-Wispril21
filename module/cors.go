package module

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewCORSModule(g *gin.Engine, devmode bool) {

	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://wisprilitb.com", "https://staging.wisprilitb.com"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}

	if (devmode){
		corsConfig.AllowAllOrigins = true
		corsConfig.AllowOrigins = nil
	}
	g.Use(cors.New(corsConfig))
}
