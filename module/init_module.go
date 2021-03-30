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

type MessageModule struct {
	usecase    entity.MessageUsecase
	controller entity.MessageController
	repo       entity.MessageRepository
}

func NewJurusanModule(g *gin.Engine) JurusanModule {
	db := database.MysqlConnect()
	jurusanRepository := repository.NewJurusanRepository(db)
	jurusanUsecase := usecase.NewJurusanUsecase(jurusanRepository)
	jurusanController := controller.NewJurusanController(g, jurusanUsecase)

	if db != nil {
		db.AutoMigrate(&entity.Jurusan{})
		db.AutoMigrate(&entity.Wisudawan{})
		db.AutoMigrate(&entity.Message{})
	}

	return JurusanModule{
		controller: jurusanController,
		usecase:    jurusanUsecase,
		repo:       jurusanRepository,
	}
}

func NewMessageModule(g *gin.Engine) MessageModule {
	db := database.MysqlConnect()

	messageRepository := repository.NewMessageRepository(db)
	messageUsecase := usecase.NewMessageUsecase(messageRepository)
	messageController := controller.NewMessageController(g, messageUsecase)
	if db != nil {
		db.AutoMigrate(&entity.Jurusan{})
		db.AutoMigrate(&entity.Wisudawan{})
		db.AutoMigrate(&entity.Message{})
	}

	return MessageModule{
		controller: messageController,
		usecase:    messageUsecase,
		repo:       messageRepository,
	}
}
