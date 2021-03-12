package entity

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paradewisudaitb/Backend/common/domain"
)

type Wisudawan struct {
	domain.EntityBase
	Nim          int
	Nama         string `gorm:"type:VARCHAR(255);not null"`
	Panggilan    string `gorm:"type:VARCHAR(255);not null"`
	JudulTA      string `gorm:"type:VARCHAR(255);not null"`
	Angkatan     int
	Jurusan      string
	Instagram    string `gorm:"type:VARCHAR(255)"`
	Linkedin     string `gorm:"type:VARCHAR(255)"`
	Twitter      string `gorm:"type:VARCHAR(255)"`
	TempatLahir  string `gorm:"type:VARCHAR(255)"`
	TanggalLahir time.Time
	Photo        string `gorm:"type:VARCHAR(255)"`
}

type LiteWisudawanSerializer struct {
	Nim  string
	Nama string
}

type FullWisudawanSerializer struct {
	LiteWisudawanSerializer
	// yang blom ada
}

type CreateWisudawanSerializer struct {
	Nim          int        `json:"nim"`
	Nama         string     `json:"nama"`
	Panggilan    string     `json:"nama_panggilan"`
	JudulTA      string     `json:"judul_ta"`
	Angkatan     int        `json:"angkatan"`
	Jurusan      string     `json:"jurusan"`
	Instagram    *string    `json:"instagram,omitempty"`
	Linkedin     *string    `json:"linkedin,omitempty"`
	Twitter      *string    `json:"twitter,omitempty"`
	TempatLahir  *string    `json:"tempat_lahir,omitempty"`
	TanggalLahir *time.Time `json:"tanggal_lahir,omitempty"`
	Photo        *string    `json:"photo,omitempty"`
}

type UpdateWisudawanSerializer struct {
	IdWisudawan  string     `json:"id_wisudawan"`
	NIM          *int       `json:"nim,omitempty"`
	Nama         *string    `json:"nama,omitempty"`
	Panggilan    *string    `json:"nama_panggilan,omitempty"`
	JudulTA      *string    `json:"judul_ta,omitempty"`
	Angkatan     *int       `json:"angkatan,omitempty"`
	Jurusan      *string    `json:"jurusan,omitempty"`
	Instagram    *string    `json:"instagram,omitempty"`
	Linkedin     *string    `json:"linkedin,omitempty"`
	Twitter      *string    `json:"twitter,omitempty"`
	TempatLahir  *string    `json:"tempat_lahir,omitempty"`
	TanggalLahir *time.Time `json:"tanggal_lahir,omitempty"`
	Photo        *string    `json:"photo,omitempty"`
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
	GetWisudawan(idWisudawan string) error
	GetAllWisudawan(Wisudawan string) error //
}

type WisudawanRepository interface {
	GetOne()
	GetAll()
	AddOne(nim, angkatan int, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo string, tanggal_lahir time.Time)
	UpdateOne(id_wisudawan string, nim, angkatan *int, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo *string, tanggal_lahir *time.Time)
	DeleteOne(id_wisudawan string)
}
