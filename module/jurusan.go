package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
	"gorm.io/gorm"
)

type JurusanModule struct {
	controller entity.JurusanController
	usecase    entity.JurusanUseCase
	repo       entity.JurusanRepository
}

func NewJurusanModule(db *gorm.DB, g *gin.Engine) JurusanModule {
	jurusanRepository := repository.NewJurusanRepository(db)
	jurusanUsecase := usecase.NewJurusanUsecase(jurusanRepository)
	jurusanController := controller.NewJurusanController(g, jurusanUsecase)

	if db != nil {
		db.AutoMigrate(&entity.Jurusan{})
	}

	return JurusanModule{
		controller: jurusanController,
		usecase:    jurusanUsecase,
		repo:       jurusanRepository,
	}
}
