package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
	"gorm.io/gorm"
)

type ContentModule struct {
	controller entity.ContentController
	usecase    entity.ContentUseCase
	repo       entity.ContentRepository
}

func NewContentModule(db *gorm.DB, g *gin.Engine) ContentModule {
	contentRepository := repository.NewContentRepository(db)
	contentUsecase := usecase.NewContentUsecase(contentRepository)
	contentController := controller.NewContentController(g, contentUsecase)

	if db != nil {
		db.AutoMigrate(&entity.Content{})
		if (!db.Migrator().HasConstraint(&entity.Content{}, "Wisudawan")) {
			db.Migrator().CreateConstraint(&entity.Content{}, "Wisudawan")
		}
	}

	return ContentModule{
		controller: contentController,
		usecase:    contentUsecase,
		repo:       contentRepository,
	}
}

func ResetContent(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entity.Content{})
	}
}
