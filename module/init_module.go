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

type MessageModule struct {
	controller entity.MessageController
	usecase    entity.MessageUsecase
	repo       entity.MessageRepository
}

type WisudawanModule struct {
	controller entity.WisudawanController
	usecase    entity.WisudawanUsecase
	repo       entity.WisudawanRepository
}

func Init(db *gorm.DB, g *gin.Engine) {
	NewJurusanModule(db, g)
	NewWisudawanModule(db, g)
	NewMessageModule(db, g)
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

func NewWisudawanModule(db *gorm.DB, g *gin.Engine) WisudawanModule {
	wisudawanRepository := repository.NewWisudawanRepository(db)
	wisudawanUsecase := usecase.NewWisudawanUsecase(wisudawanRepository)
	wisudawanController := controller.NewWisudawanController(g, wisudawanUsecase)

	if db != nil {
		db.AutoMigrate(&entity.Wisudawan{})
	}
	return WisudawanModule{
		controller: wisudawanController,
		usecase:    wisudawanUsecase,
		repo:       wisudawanRepository,
	}
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
