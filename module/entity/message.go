package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
)

type Message struct {
	domain.EntityBase
	ReceiverID string
	Receiver   Wisudawan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message    string    `gorm:"type:text;"`
	Sender     string    `gorm:"type:VARCHAR(255);not null"`
}

type CreateMessageSerializer struct {
	IdWisudawan string `json:"id_wisudawan" binding:"required"`
	Message     string `json:"message" binding:"required"`
	Sender      string `json:"sender" binding:"required;lte=255"`
}

type MessageController interface {
	CreateMessage(gin.Context) error
	DeleteMessage(gin.Context) error
}

type MessageUsecase interface {
	CreateMessage(item CreateMessageSerializer) error
	DeleteMessage(idMessage string) error
}

type MessageRepository interface {
	AddOne(message, sender, idWisudawan string) error
	DeleteOne(idMessage string) error
}
