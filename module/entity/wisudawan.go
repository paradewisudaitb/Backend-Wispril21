package entity

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
	uuid "github.com/satori/go.uuid"
)

type Wisudawan struct {
	domain.EntityBase
	Nim           uint32
	Nama          string `gorm:"type:VARCHAR(255);not null"`
	Panggilan     string `gorm:"type:VARCHAR(255);not null"`
	JudulTA       string `gorm:"type:VARCHAR(255);not null"`
	Angkatan      uint16 `gorm:"type:SMALLINT;not null"`
	JurusanID     string
	Jurusan       Jurusan `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Instagram     string  `gorm:"type:VARCHAR(255)"`
	Linkedin      string  `gorm:"type:VARCHAR(255)"`
	Twitter       string  `gorm:"type:VARCHAR(255)"`
	TempatLahir   string  `gorm:"type:VARCHAR(255)"`
	TanggalLahir  time.Time
	Photo         string  `gorm:"type:VARCHAR(255)"`
	Organizations []*Orgz `gorm:"many2many:wisudawan_orgz;"`
}

func (Wisudawan) TableName() string {
	return "wisudawan"
}

type SimpleWisudawanSerializer struct {
	Nim  uint32
	Nama string
}

type CreateWisudawanSerializer struct {
	Nim          uint32    `json:"nim"`
	Nama         string    `json:"nama"`
	Panggilan    string    `json:"nama_panggilan"`
	JudulTA      string    `json:"judul_ta"`
	Angkatan     uint16    `json:"angkatan"`
	Jurusan      string    `json:"jurusan"`
	Instagram    string    `json:"instagram"`
	Linkedin     string    `json:"linkedin"`
	Twitter      string    `json:"twitter"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Photo        string    `json:"photo"`
}

type UpdateWisudawanSerializer struct {
	IdWisudawan  string    `json:"id_wisudawan"`
	NIM          uint32    `json:"nim"`
	Nama         string    `json:"nama"`
	Panggilan    string    `json:"nama_panggilan"`
	JudulTA      string    `json:"judul_ta"`
	Angkatan     uint16    `json:"angkatan"`
	Jurusan      string    `json:"jurusan"`
	Instagram    string    `json:"instagram"`
	Linkedin     string    `json:"linkedin"`
	Twitter      string    `json:"twitter"`
	TempatLahir  string    `json:"tempat_lahir"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	Photo        string    `json:"photo"`
}

type DeleteWisudawanSerializer struct {
	IdWisudawan string `json:"id_wisudawan"`
}

type WisudawanController interface {
	CreateWisudawan(gin.Context) error
	UpdateWisudawan(gin.Context) error
	DeleteWisudawan(gin.Context) error
	GetWisudawan(gin.Context) error
}

type WisudawanUsecase interface {
	CreateWisudawan(item CreateWisudawanSerializer) error
	DeleteWisudawan(item DeleteWisudawanSerializer) error
	UpdateWisudawan(item UpdateWisudawanSerializer) error
	GetWisudawan(idWisudawan string) (Wisudawan, error)
	GetAllWisudawan(wisudawan string) ([]Wisudawan, error)
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
