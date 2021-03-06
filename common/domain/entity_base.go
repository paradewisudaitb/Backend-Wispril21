package domain

import (
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type EntityBase struct {
	ID        string     `gorm:"primary_key;" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (EntityBase) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	scope.SetColumn("id", uuid)
	return nil
}
