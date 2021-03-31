package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
	"gorm.io/gorm"
)

type OrgzModule struct {
	controller entity.OrgzController
	usecase    entity.OrgzUseCase
	repo       entity.OrgzRepository
}

func NewOrgzModule(db *gorm.DB, g *gin.Engine) OrgzModule {
	orgzRepository := repository.NewOrgzRepository(db)
	orgzUsecase := usecase.NewOrgzUsecase(orgzRepository)
	orgzController := controller.NewOrgzController(g, orgzUsecase)
	if db != nil {
		db.AutoMigrate(&entity.Orgz{})
	}

	return OrgzModule{
		controller: orgzController,
		usecase:    orgzUsecase,
		repo:       orgzRepository,
	}
}

func ResetOrgz(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entity.Orgz{})
	}
}
