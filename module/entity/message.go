package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
)

type Message struct {
	domain.EntityBase
	//ID Wisudawan
	message string
	sender  string
}

type CreateMessageSerializer struct {
	IdWisudawan string `json:"id_wisudawan"`
	Message     string `json:"message"`
	Sender      string `json:"sender"`
}

type DeleteMessageSerializer struct {
	IdWisudawan string `json:"id_wisudawan"`
	IdMessage   string `json:"id_message"`
}

type MessageController interface {
	CreateMessage(gin.Context) error
	DeleteMessage(gin.Context) error
}

type MessageUsecase interface {
	CreateMessage(item CreateMessageSerializer) error
	DeleteMessage(item DeleteMessageSerializer) error
}

type MessageRepository interface {
	AddOne(message, sender, idWisudawan string) error
	DeleteOne(idMessage, idWisudawan string) error
}
