package entity

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Wisudawan struct {
	domain.EntityBase
	Nim          uint32    `json:"nim"`
	Nama         string    `json:"nama" gorm:"type:VARCHAR(255);not null"`
	Panggilan    string    `json:"nama_panggilan" gorm:"type:VARCHAR(255);not null"`
	JudulTA      string    `json:"judul_ta" gorm:"type:VARCHAR(255);not null"`
	Angkatan     uint16    `json:"angkatan" gorm:"type:SMALLINT;not null"`
	JurusanID    string    `json:"jurusan_id" gorm:"type:VARCHAR(50)"`
	Jurusan      Jurusan   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Instagram    string    `json:"instagram" gorm:"type:VARCHAR(255)"`
	Linkedin     string    `json:"linkedin" gorm:"type:VARCHAR(255)"`
	Twitter      string    `json:"twitter" gorm:"type:VARCHAR(255)"`
	TempatLahir  string    `json:"tempat_lahir" gorm:"type:VARCHAR(255)"`
	TanggalLahir time.Time `json:"tanggal_lahir" gorm:"type:DATE"`
	Photo        string    `json:"photo" gorm:"type:VARCHAR(255)"`
}

func (Wisudawan) TableName() string {
	return "wisudawan"
}

type SimpleWisudawanSerializer struct {
	Nim  uint32
	Nama string
}

type CreateWisudawanSerializer struct {
	Nim          uint32    `json:"nim" wispril:"required"`
	Nama         string    `json:"nama" wispril:"required" binding:"lte=255"`
	Panggilan    string    `json:"nama_panggilan" wispril:"required" binding:"lte=255"`
	JudulTA      string    `json:"judul_ta" wispril:"required" binding:"lte=255"`
	Angkatan     uint16    `json:"angkatan" wispril:"required" binding:"lte=25"`
	Jurusan      string    `json:"jurusan_id"`
	Instagram    string    `json:"instagram" binding:"lte=255"`
	Linkedin     string    `json:"linkedin" binding:"lte=255"`
	Twitter      string    `json:"twitter" binding:"lte=255"`
	TempatLahir  string    `json:"tempat_lahir" binding:"lte=255"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Photo        string    `json:"photo" binding:"lte=255"`
}

type UpdateWisudawanSerializer struct {
	IdWisudawan  uuid.UUID `json:"id_wisudawan"`
	NIM          uint32    `json:"nim" wispril:"required"`
	Nama         string    `json:"nama" binding:"lte=255"`
	Panggilan    string    `json:"nama_panggilan" binding:"lte=255"`
	JudulTA      string    `json:"judul_ta" binding:"lte=255"`
	Angkatan     uint16    `json:"angkatan" binding:"lte=25"`
	Jurusan      string    `json:"jurusan_id"`
	Instagram    string    `json:"instagram" binding:"lte=255"`
	Linkedin     string    `json:"linkedin" binding:"lte=255"`
	Twitter      string    `json:"twitter" binding:"lte=255"`
	TempatLahir  string    `json:"tempat_lahir"  binding:"lte=255"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Photo        string    `json:"photo"  binding:"lte=255"`
}

type WisudawanController interface {
	CreateWisudawan(ctx *gin.Context)
	UpdateWisudawan(ctx *gin.Context)
	DeleteWisudawan(ctx *gin.Context)
	GetWisudawan(ctx *gin.Context)
}

type WisudawanUsecase interface {
	CreateWisudawan(item CreateWisudawanSerializer) error
	DeleteWisudawan(idWisudawan uuid.UUID) error
	UpdateWisudawan(item UpdateWisudawanSerializer) error
	GetWisudawan(idWisudawan uuid.UUID) (Wisudawan, error)
	GetAllWisudawan() ([]Wisudawan, error)
	FilterWisudawan(jurusan string) ([]Wisudawan, error)
}

type WisudawanRepository interface {
	GetOne(wisudawanID uuid.UUID) (Wisudawan, error)
	GetAll() ([]Wisudawan, error)
	AddOne(nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusan, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error
	UpdateOne(WisudawanID uuid.UUID, nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusanID, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error
	DeleteOne(WisudawanID uuid.UUID) error
	// Filter(jurusan string) ([]Wisudawan, error)
}
