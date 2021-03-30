package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	domain.EntityBase
	ReceiverID string    `gorm:"type:VARCHAR(50)" json:"id_wisudawan"`
	Receiver   Wisudawan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Message    string    `gorm:"type:TEXT;" json:"message"`
	Sender     string    `gorm:"type:VARCHAR(255);not null" json:"sender"`
}

type CreateMessageSerializer struct {
	IdWisudawan string `json:"id_wisudawan" binding:"required"`
	Message     string `json:"message" binding:"required"`
	Sender      string `json:"sender" binding:"required,lte=255"`
}

type MessageController interface {
	CreateMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
	GetMessage(ctx *gin.Context)
}

type MessageUsecase interface {
	CreateMessage(item CreateMessageSerializer) error
	DeleteMessage(idMessage uuid.UUID) error
	GetMessage(idWisudawan uuid.UUID) ([]Message, error)
}

type MessageRepository interface {
	AddOne(message, sender, idWisudawan string) error
	DeleteOne(idMessage string) error
	GetMessage(idWisudawan string) ([]Message, error)
}
