package entity

import (
	"github.com/gin-gonic/gin"
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
	Name             string `json:"name" wispril:"required" binding:"lte=255"`
	Category         string `json:"category" wispril:"required" binding:"lte=64"`
	Logo             string `json:"logo" wispril:"required" binding:"lte=255"`
	ApresiasiPoster  string `json:"apresiasi_poster" binding:"lte=255"`
	ApresiasiTulisan string `json:"apresiasi_tulisan"`
	ApresiasiVideo   string `json:"apresiasi_video" binding:"lte=255"`
}

type UpdateOrgzSerializer struct {
	IdOrgz           string `json:"id_organization" wispril:"required"`
	Name             string `json:"name" binding:"lte=255"`
	Category         string `json:"category" binding:"lte=64"`
	Logo             string `json:"logo" binding:"lte=255"`
	ApresiasiPoster  string `json:"apresiasi_poster" binding:"lte=255"`
	ApresiasiTulisan string `json:"apresiasi_tulisan"`
	ApresiasiVideo   string `json:"apresiasi_video" binding:"lte=255"`
}

type OrgzController interface {
	CreateOrgz(ctx *gin.Context)
	UpdateOrgz(ctx *gin.Context)
	DeleteOrgz(ctx *gin.Context)
	GetOrgz(ctx *gin.Context)
}

type OrgzUseCase interface {
	CreateOrgz(item CreateOrgzSerializer) error
	DeleteOrgz(idOrgz string) error
	UpdateOrgz(item UpdateOrgzSerializer) error
	GetOrgz(idOrgz string) (Orgz, error)
}

type OrgzRepository interface {
	GetOne(idOrgz string) (Orgz, error)
	AddOne(name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string) error
	UpdateOne(idOrgz string, name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string) error
	DeleteOne(idOrgz string) error
}
