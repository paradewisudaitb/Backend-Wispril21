package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
)

func Init(db *gorm.DB, g *gin.Engine) {
	NewJurusanModule(db, g)
	NewWisudawanModule(db, g)
	NewMessageModule(db, g)
	db.AutoMigrate(&entity.Orgz{})
	db.AutoMigrate(&entity.Content{})
}
