package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/connection/database"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
)

type JurusanModule struct {
	usecase    entity.JurusanUseCase
	controller entity.JurusanController
	repo       entity.JurusanRepository
}

func NewJurusanModule(g *gin.Engine) JurusanModule {
	db := database.PostgresConnect()
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
