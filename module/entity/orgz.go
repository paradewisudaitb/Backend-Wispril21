package entity

import (
	"github.com/paradewisudaitb/Backend/common/domain"
)

type Orgz struct {
	domain.EntityBase
	Name             string `gorm:"type:VARCHAR(255);not null"`
	Category         string `gorm:"type:VARCHAR(64);not null"`
	Logo             string `gorm:"type:VARCHAR(255);not null"`
	ApresiasiPoster  string `gorm:"type:VARCHAR(255);"`
	ApresiasiTulisan string `gorm:"type:text;"`
	ApresiasiVideo   string `gorm:"type:VARCHAR(255);"`
}

type CreateOrgzSerializer struct {
	Name             string  `json:"name"`
	Category         string  `json:"category"`
	Logo             string  `json:"logo"`
	ApresiasiPoster  *string `json:"apresiasi_poster"`
	ApresiasiTulisan *string `json:"apresiasi_tulisan"`
	ApresiasiVideo   *string `json:"apresiasi_video"`
}

type UpdateOrgzSerializer struct {
	IdOrgz           string  `json:"id_organization"`
	Name             *string `json:"name"`
	Category         *string `json:"category"`
	Logo             *string `json:"logo"`
	ApresiasiPoster  *string `json:"apresiasi_poster"`
	ApresiasiTulisan *string `json:"apresiasi_tulisan"`
	ApresiasiVideo   *string `json:"apresiasi_video"`
}

type DeleteOrgzSerializer struct {
	IdOrgz string `json:"id_organization"`
}

type OrgzUseCase interface {
	CreateOrgz(item CreateOrgzSerializer) error
	DeleteOrgz(item DeleteOrgzSerializer) error
	UpdateOrgz(item UpdateOrgzSerializer) error
	GetOrgz(idOrgz string) (Orgz, error)
}

type OrgzRepository interface {
	GetOne(id_orgz string) (Orgz, error)
	AddOne(name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string)
	UpdateOne(id_organization string, name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video *string) error
	DeleteOne(id_orgz string) error
}
