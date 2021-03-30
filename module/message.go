package module

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/module/controller"
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	"github.com/paradewisudaitb/Backend/module/usecase"
	"gorm.io/gorm"
)

type MessageModule struct {
	controller entity.MessageController
	usecase    entity.MessageUsecase
	repo       entity.MessageRepository
}

func NewMessageModule(db *gorm.DB, g *gin.Engine) MessageModule {
	messageRepository := repository.NewMessageRepository(db)
	messageUsecase := usecase.NewMessageUsecase(messageRepository)
	messageController := controller.NewMessageController(g, messageUsecase)
	if db != nil {
		db.AutoMigrate(&entity.Message{})
	}

	return MessageModule{
		controller: messageController,
		usecase:    messageUsecase,
		repo:       messageRepository,
	}
}

func ResetMessage(db *gorm.DB) {
	if db != nil {
		db.Migrator().DropTable(&entity.Message{})
	}
}
