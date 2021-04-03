package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
	"gorm.io/gorm"
)

type WisudawanModule struct {
	controller entity.WisudawanController
	usecase    entity.WisudawanUsecase
	repo       entity.WisudawanRepository
}

func NewWisudawanModule(db *gorm.DB, g *gin.Engine) WisudawanModule {
	wisudawanRepository := repository.NewWisudawanRepository(db)
	wisudawanUsecase := usecase.NewWisudawanUsecase(wisudawanRepository)
	wisudawanController := controller.NewWisudawanController(g, wisudawanUsecase)

	if db != nil {
		db.AutoMigrate(&entity.Wisudawan{})
		if (!db.Migrator().HasConstraint(&entity.Wisudawan{}, "Jurusan")) {
			db.Migrator().CreateConstraint(&entity.Wisudawan{}, "Jurusan")
		}
	}
	return WisudawanModule{
		controller: wisudawanController,
		usecase:    wisudawanUsecase,
		repo:       wisudawanRepository,
	}
}

func ResetWisudawan(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entity.Wisudawan{})
	}
}
