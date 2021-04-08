package module

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, g *gin.Engine, devmode bool) {
	NewCORSModule(g, devmode)
	NewLimiterModule(g)
	NewJurusanModule(db, g)
	NewWisudawanModule(db, g)
	NewMessageModule(db, g)
	NewOrgzModule(db, g)
	NewContentModule(db, g)
	g.GET("/reset", middleware.ResetAuth, func(c *gin.Context) {
		Reset(db, g)
	})
	g.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong!")
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
