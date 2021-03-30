package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	domain.EntityBase
	ReceiverID string    `gorm:"type:VARCHAR(50);not null" json:"id_wisudawan"`
	Receiver   Wisudawan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Message    string    `gorm:"type:TEXT;not null" json:"message"`
	Sender     string    `gorm:"type:VARCHAR(255);not null" json:"sender"`
}

func (Message) TableName() string {
	return "message"
}

type CreateMessageSerializer struct {
	IdWisudawan string `json:"id_wisudawan" wispril:"required"`
	Message     string `json:"message" wispril:"required"`
	Sender      string `json:"sender" wispril:"required" binding:"lte=255"`
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
