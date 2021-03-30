package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Jurusan struct {
	domain.EntityBase
	Jurusan       string `gorm:"type:VARCHAR(50);not null" json:"jurusan"`
	Fakultas      string `gorm:"type:VARCHAR(50);not null" json:"fakultas"`
	FakultasShort string `gorm:"type:VARCHAR(5);not null" json:"fakultas_short"`
	JurusanShort  string `gorm:"type:VARCHAR(5);not null" json:"jurusan_short"`
}

func (Jurusan) TableName() string {
	return "jurusan"
}

type CreateJurusanSerializer struct {
	Jurusan       string `json:"jurusan" wispril:"required"`
	Fakultas      string `json:"fakultas" wispril:"required"`
	FakultasShort string `json:"fakultas_short" wispril:"required" binding:"lte=5"`
	JurusanShort  string `json:"jurusan_short" wispril:"required" binding:"lte=5"`
}

type UpdateJurusanSerializer struct {
	IdJurusan     uuid.UUID `json:"id_jurusan" wispril:"required"`
	Jurusan       string    `json:"jurusan"`
	Fakultas      string    `json:"fakultas"`
	FakultasShort string    `json:"fakultas_short" binding:"lte=5"`
	JurusanShort  string    `json:"jurusan_short" binding:"lte=5"`
}

type JurusanController interface {
	CreateJurusan(ctx *gin.Context)
	UpdateJurusan(ctx *gin.Context)
	DeleteJurusan(ctx *gin.Context)
	GetJurusan(ctx *gin.Context)
}

type JurusanUseCase interface {
	CreateJurusan(item CreateJurusanSerializer) error
	DeleteJurusan(IdJurusan uuid.UUID) error
	UpdateJurusan(item UpdateJurusanSerializer) error
	GetJurusan(IdJurusan uuid.UUID) (Jurusan, error)
}

type JurusanRepository interface {
	GetOne(id uuid.UUID) (Jurusan, error)
	AddOne(jurusan, fakultas, fakultas_short, jurusan_short string) error
	UpdateOne(id uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short string) error
	DeleteOne(id uuid.UUID) error
}
