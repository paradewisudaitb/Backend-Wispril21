// entity
// entitybase, id_wisudawan, ip, time

// Usecase AddViews

// Repository AddViews

// request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record
package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type View struct {
	ID          string    `gorm:"primaryKey;type:VARCHAR(50)" json:"id"`
	WisudawanID string    `gorm:"type:VARCHAR(50);not null" json:"id_wisudawan"`
	IP          string    `gorm:"type:VARCHAR(255);not null" json:"ip"`
	AccessTime  time.Time `json:"time"`
}

type GetViewWisudawan struct {
	Wisudawan Wisudawan `gorm:"embedded" json:"wisudawan"`
	Count     int64     `json:"count"`
}

func (e *View) BeforeCreate(scope *gorm.DB) error {
	e.ID = uuid.NewV4().String()
	return nil
}

func (View) TableName() string {
	return "view"
}

type ViewController interface {
	AddView(ctx *gin.Context)
}

type ViewUseCase interface {
	AddView(idWisudawan uuid.UUID, clientIP string) error
	GetTop5() ([]GetViewWisudawan, error)
}

type ViewRepository interface {
	AddOne(IdWisudawan string, IP string, Time time.Time) error
	GetLast(IdWisudawan string, IP string) (View, error)
	GetTop5() ([]GetViewWisudawan, error)
}
