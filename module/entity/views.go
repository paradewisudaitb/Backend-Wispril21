// entity
// entitybase, id_wisudawan, ip, time

// Usecase AddViews

// Repository AddViews

// request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record
// package entity

// import (
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/paradewisudaitb/Backend/common/domain"
// 	uuid "github.com/satori/go.uuid"
// )

// type Views struct {
// 	domain.EntityBase
// 	IdWisudawan uuid.UUID `gorm:"type:VARCHAR(255);not null" json:"id_wisudawan"`
// 	IP          string    `gorm:"type:VARCHAR(255);not null" json:"ip"`
// 	Time        time.Time `gorm:"type:DATE" json:"time"`
// }

// func (Views) TableName() string {
// 	return "views"
// }

// type ViewsController interface {
// 	AddViews(ctx *gin.Context)
// }

// type ViewsUseCase interface {
// 	AddViews(IP string) (Views, error)
// }

// type ViewsRepository interface {
// 	AddOne(IdWisudawan uuid.UUID, IP string, Time time.Time) error
// }
