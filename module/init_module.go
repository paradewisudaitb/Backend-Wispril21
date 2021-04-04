package module

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, g *gin.Engine) {
	g.Use(limit.NewRateLimiter(func(c *gin.Context) string {
		return c.ClientIP() // limit rate by client ip
	}, func(c *gin.Context) (*rate.Limiter, time.Duration) {
		return rate.NewLimiter(rate.Every(500*time.Millisecond), 4), time.Hour
	}, func(c *gin.Context) {
		if values := c.Request.Header.Get("Authorization"); len(values) > 0 {
			token := os.Getenv("AUTH_TOKEN")
			if values == token {
				c.Next()
				return
			}
		}
		controller.ForceResponse(c, http.StatusTooManyRequests, "too_many_requests")
	}))

	NewJurusanModule(db, g)
	NewWisudawanModule(db, g)
	NewMessageModule(db, g)
	NewOrgzModule(db, g)
	NewContentModule(db, g)
	db.AutoMigrate(&entity.View{})
	g.GET("/reset", middleware.ResetAuth, func(c *gin.Context) {
		Reset(db, g)
	})
}

func Reset(db *gorm.DB, g *gin.Engine) {
	ResetJurusan(db)
	ResetWisudawan(db)
	ResetMessage(db)
	ResetOrgz(db)
	ResetContent(db)
	os.Exit(0)
}
