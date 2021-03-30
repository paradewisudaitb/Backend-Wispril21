package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Content struct {
	domain.EntityBase
	WisudawanID    string    `json:"id_wisudawan" gorm:"type:VARCHAR(50);not null"`
	Wisudawan      Wisudawan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	OrganizationID string    `json:"id_organization" gorm:"type:VARCHAR(50)"`
	Organization   Orgz      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Type           string    `gorm:"type:VARCHAR(16);not null" json:"content_type"`
	Headings       string    `gorm:"type:VARCHAR(255);not null" json:"headings"`
	Details        string    `gorm:"type:TEXT" json:"details"`
	Image          string    `gorm:"type:VARCHAR(255)" json:"image"`
}

func (Content) TableName() string {
	return "contents"
}

type CreateContentSerializer struct {
	Wisudawan    string `json:"id_wisudawan" wispril:"required"`
	Organization string `json:"id_organization"`
	ContentType  string `json:"content_type" wispril:"required" binding:"lte=16"`
	Headings     string `json:"headings" wispril:"required" binding:"lte=255"`
	Details      string `json:"details"`
	Image        string `json:"image" binding:"lte=255"`
}

type UpdateContentSerializer struct {
	Content      string `json:"id_content" wispril:"required"`
	Wisudawan    string `json:"id_wisudawan"`
	Organization string `json:"id_organization"`
	ContentType  string `json:"content_type" binding:"lte=16"`
	Headings     string `json:"headings" binding:"lte=255"`
	Details      string `json:"details"`
	Image        string `json:"image" binding:"lte=255"`
}

type ContentController interface {
	CreateContent(ctx *gin.Context)
	UpdateContent(ctx *gin.Context)
	DeleteContent(ctx *gin.Context)
	GetContent(ctx *gin.Context)
	GetContentByWisudawan(ctx *gin.Context)
}

type ContentUseCase interface {
	CreateContent(item CreateContentSerializer) error
	DeleteContent(IdContent uuid.UUID) error
	UpdateContent(item UpdateContentSerializer) error
	GetContent(IdContent uuid.UUID) (Content, error)
	GetByWisudawan(IdWisudawan uuid.UUID) ([]Content, error)
}

type ContentRepository interface {
	GetOne(id string) (Content, error)
	GetByWisudawan(idWisudawan string) ([]Content, error)
	AddOne(idWisudawan, idOrgz, contenttype, headings, details, image string) error
	UpdateOne(idContent string, idWisudawan, idOrgz, contenttype, headings, details, image string) error
	DeleteOne(idContent string) error
}
