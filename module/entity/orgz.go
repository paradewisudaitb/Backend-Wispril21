package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Orgz struct {
	domain.EntityBase
	Name             string `gorm:"type:VARCHAR(255);not null" json:"name"`
	Slug             string `gorm:"type:VARCHAR(128);not null;unique" json:"slug"`
	Category         string `gorm:"type:VARCHAR(64);not null" json:"category"`
	Logo             string `gorm:"type:VARCHAR(255);not null" json:"logo"`
	ApresiasiPoster  string `gorm:"type:VARCHAR(255);" json:"apresiasi_poster"`
	ApresiasiTulisan string `gorm:"type:text;" json:"apresiasi_tulisan"`
	ApresiasiVideo   string `gorm:"type:VARCHAR(255);" json:"apresiasi_video"`
	FakultasShort    string `gorm:"type:VARCHAR(5)" json:"fakultas_short"`
}

func (Orgz) TableName() string {
	return "organization"
}

type CreateOrgzSerializer struct {
	Name             string `json:"name" wispril:"required" binding:"lte=255"`
	Slug             string `json:"slug" wispril:"required" binding:"lte=255"`
	Category         string `json:"category" wispril:"required" binding:"lte=64"`
	Logo             string `json:"logo" wispril:"required" binding:"lte=255"`
	ApresiasiPoster  string `json:"apresiasi_poster" binding:"lte=255"`
	ApresiasiTulisan string `json:"apresiasi_tulisan"`
	ApresiasiVideo   string `json:"apresiasi_video" binding:"lte=255"`
	FakultasShort    string `json:"fakultas_short" binding:"lte=5"`
}

type UpdateOrgzSerializer struct {
	IdOrgz           string `json:"id_organization" wispril:"required"`
	Slug             string `json:"slug" binding:"lte=255"`
	Name             string `json:"name" binding:"lte=255"`
	Category         string `json:"category" binding:"lte=64"`
	Logo             string `json:"logo" binding:"lte=255"`
	ApresiasiPoster  string `json:"apresiasi_poster" binding:"lte=255"`
	ApresiasiTulisan string `json:"apresiasi_tulisan"`
	ApresiasiVideo   string `json:"apresiasi_video" binding:"lte=255"`
	FakultasShort    string `json:"fakultas_short" binding:"lte=5"`
}

type OrgzController interface {
	CreateOrgz(ctx *gin.Context)
	UpdateOrgz(ctx *gin.Context)
	DeleteOrgz(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetBySlug(ctx *gin.Context)
	GetAll(ctx *gin.Context)
}

type OrgzUseCase interface {
	CreateOrgz(item CreateOrgzSerializer) error
	DeleteOrgz(idOrgz uuid.UUID) error
	UpdateOrgz(item UpdateOrgzSerializer) error
	GetOrgz(idOrgz uuid.UUID) (Orgz, error)
	GetAll() ([]Orgz, error)
	GetBySlug(slug string) (Orgz, error)
}

type OrgzRepository interface {
	GetOne(idOrgz string) (Orgz, error)
	AddOne(name, slug, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video, fakultas_short string) error
	UpdateOne(idOrgz, name, slug, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video, fakultas_short string) error
	DeleteOne(idOrgz string) error
	GetAll() ([]Orgz, error)
	GetBySlug(slug string) (Orgz, error)
}
