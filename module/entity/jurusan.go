package entity

import (
	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Jurusan struct {
	domain.EntityBase
	IdJurusan     uuid.UUID `gorm:"type:uuid;primary_key;"`
	Jurusan       string    `gorm:"type:VARCHAR(50);not null"`
	Fakultas      string    `gorm:"type:VARCHAR(50);not null"`
	FakultasShort string    `gorm:"type:VARCHAR(5);not null"`
	JurusanShort  string    `gorm:"type:VARCHAR(5);not null"`
}

type CreateJurusanSerializer struct {
	IdJurusan     uuid.UUID `json:"id_wisudawan"`
	Jurusan       string    `json:"jurusan"`
	Fakultas      string    `json:"fakultas"`
	FakultasShort string    `json:"fakultas_short"`
	JurusanShort  string    `json:"jurusan_short"`
}

type UpdateJurusanSerializer struct {
	IdJurusan     uuid.UUID `json:"id_wisudawan"`
	Jurusan       *string   `json:"jurusan"`
	Fakultas      *string   `json:"fakultas"`
	FakultasShort *string   `json:"fakultas_short"`
	JurusanShort  *string   `json:"jurusan_short"`
}

type DeleteJurusanSerializer struct {
	IdJurusan uuid.UUID `json:"id_wisudawan"`
}

type JurusanController interface {
	CreateJurusan(gin.Context) error
	UpdateJurusan(gin.Context) error
	DeleteJurusan(gin.Context) error
	GetJurusan(gin.Context) error
}

type JurusanUsecase interface {
	CreateJurusan(item CreateJurusanSerializer) error
	DeleteJurusan(item DeleteJurusanSerializer) error
	UpdateJurusan(item UpdateJurusanSerializer) error
	GetJurusan(idJurusan string) error
	GetAllJurusan(Jurusan string) error
}

type JurusanRepository interface {
	AddOne(idJurusan, Jurusan, Fakultas, FakultasShort, JurusanShort string) error
	UpdateOne(idJurusan, Jurusan, Fakultas, FakultasShort, JurusanShort *string) error
	DeleteOne(idJurusan string) error
}
