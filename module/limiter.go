package module

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	limit "github.com/yangxikun/gin-limit-by-key"
	"golang.org/x/time/rate"
)

func NewLimiterModule(g *gin.Engine) {
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
}
