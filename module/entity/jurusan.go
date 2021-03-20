package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Jurusan struct {
	domain.EntityBase
	Jurusan       string `gorm:"type:VARCHAR(50);not null"`
	Fakultas      string `gorm:"type:VARCHAR(50);not null"`
	FakultasShort string `gorm:"type:VARCHAR(5);not null"`
	JurusanShort  string `gorm:"type:VARCHAR(5);not null"`
}

type CreateJurusanSerializer struct {
	Jurusan       *string `json:"jurusan"`
	Fakultas      *string `json:"fakultas"`
	FakultasShort *string `json:"fakultas_short"`
	JurusanShort  *string `json:"jurusan_short"`
}

type UpdateJurusanSerializer struct {
	IdJurusan     uuid.UUID `json:"id_jurusan"`
	Jurusan       *string   `json:"jurusan"`
	Fakultas      *string   `json:"fakultas"`
	FakultasShort *string   `json:"fakultas_short"`
	JurusanShort  *string   `json:"jurusan_short"`
}

type DeleteJurusanSerializer struct {
	IdJurusan uuid.UUID `json:"id_jurusan"`
}

type JurusanController interface {
	CreateJurusan(gin.Context) error
	UpdateJurusan(gin.Context) error
	DeleteJurusan(gin.Context) error
	GetJurusan(gin.Context) error
}

type JurusanUseCase interface {
	CreateJurusan(item CreateJurusanSerializer) error
	DeleteJurusan(item DeleteJurusanSerializer) error
	UpdateJurusan(item UpdateJurusanSerializer) error
	GetJurusan(IdJurusan uuid.UUID) (Jurusan, error)
}

type JurusanRepository interface {
	GetOne(id_jurusan uuid.UUID) (Jurusan, error)
	AddOne(jurusan, fakultas, fakultas_short, jurusan_short string)
	UpdateOne(id_jurusan uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short *string) error
	DeleteOne(id_jurusan uuid.UUID) error
}
