package entity

import (
	"github.com/paradewisudaitb/Backend/common/domain"
)

type Content struct {
	domain.EntityBase
	WisudawanID    string    `json:"id_wisudawan" gorm:"type:VARCHAR(50);not null"`
	Wisudawan      Wisudawan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrganizationID string    `json:"id_organization" gorm:"type:VARCHAR(50)"`
	Organization   Orgz      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Type           string    `gorm:"type:VARCHAR(16);not null" json:"content_type"`
	Headings       string    `gorm:"type:TEXT;not null" json:"headings"`
	Details        string    `gorm:"type:TEXT" json:"details"`
	Image          string    `gorm:"type:VARCHAR(255)" json:"image"`
}

func (Content) TableName() string {
	return "contents"
}

// type CreateJurusanSerializer struct {
// 	Jurusan       string `json:"jurusan" wispril:"required"`
// 	Fakultas      string `json:"fakultas" wispril:"required"`
// 	FakultasShort string `json:"fakultas_short" wispril:"required" binding:"lte=5"`
// 	JurusanShort  string `json:"jurusan_short" wispril:"required" binding:"lte=5"`
// }

// type UpdateJurusanSerializer struct {
// 	IdJurusan     uuid.UUID `json:"id_jurusan" wispril:"required"`
// 	Jurusan       string    `json:"jurusan"`
// 	Fakultas      string    `json:"fakultas"`
// 	FakultasShort string    `json:"fakultas_short" binding:"lte=5"`
// 	JurusanShort  string    `json:"jurusan_short" binding:"lte=5"`
// }

// type JurusanController interface {
// 	CreateJurusan(ctx *gin.Context)
// 	UpdateJurusan(ctx *gin.Context)
// 	DeleteJurusan(ctx *gin.Context)
// 	GetJurusan(ctx *gin.Context)
// }

// type JurusanUseCase interface {
// 	CreateJurusan(item CreateJurusanSerializer) error
// 	DeleteJurusan(IdJurusan uuid.UUID) error
// 	UpdateJurusan(item UpdateJurusanSerializer) error
// 	GetJurusan(IdJurusan uuid.UUID) (Jurusan, error)
// }

// type JurusanRepository interface {
// 	GetOne(id uuid.UUID) (Jurusan, error)
// 	AddOne(jurusan, fakultas, fakultas_short, jurusan_short string) error
// 	UpdateOne(id uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short string) error
// 	DeleteOne(id uuid.UUID) error
// }
