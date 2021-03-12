package entity

import "github.com/paradewisudaitb/Backend/common/domain"

type Orgz struct {
	domain.EntityBase
	Name             string `gorm:"type:VARCHAR(255);not null"`
	Category         string `gorm:"type:VARCHAR(64);not null"`
	Logo             string `gorm:"type:VARCHAR(255);not null"`
	ApresiasiPoster  string `gorm:"type:VARCHAR(255);"`
	ApresiasiTulisan string `gorm:"type:text;"`
	ApresiasiVideo   string `gorm:"type:VARCHAR(255);"`
}
