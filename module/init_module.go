package module

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller/middleware"
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, g *gin.Engine) {
	NewJurusanModule(db, g)
	NewWisudawanModule(db, g)
	NewMessageModule(db, g)
	db.AutoMigrate(&entity.Orgz{})
	db.AutoMigrate(&entity.Content{})
	g.GET("/reset", middleware.Auth, func(c *gin.Context) {
		Reset(db, g)
	})
}

func Reset(db *gorm.DB, g *gin.Engine) {
	ResetJurusan(db)
	ResetWisudawan(db)
	ResetMessage(db)
	db.Migrator().DropTable(&entity.Orgz{})
	db.Migrator().DropTable(&entity.Content{})
	os.Exit(0)
}
