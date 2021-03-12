package repository

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/paradewisudaitb/Backend/module/entity"
)

type WisudawanRepository struct {
	db *gorm.DB
}

func NewWisudawanRepository(db *gorm.DB) WisudawanRepository {
	return WisudawanRepository{db: db}
}

func (repo *WisudawanRepository) AddOne(nim, angkatan int, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo string, tanggal_lahir time.Time) {
	wisudawan := entity.Wisudawan{Nim: nim, Angkatan: angkatan, Nama: nama, Panggilan: panggilan, JudulTA: judul_ta, Jurusan: jurusan, Instagram: instagram, Linkedin: linkedin, Twitter: twitter, TempatLahir: tempat_lahir, Photo: photo, TanggalLahir: tanggal_lahir}
	repo.db.Create(&wisudawan)
}

func (repo *WisudawanRepository) UpdateOne(id_wisudawan string, nim, angkatan *int, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo *string, tanggal_lahir *time.Time) {
	var wisudawan entity.Wisudawan
	wisudawan_update := entity.Wisudawan{Nim: *nim, Angkatan: *angkatan, Nama: *nama, Panggilan: *panggilan, JudulTA: *judul_ta, Jurusan: *jurusan, Instagram: *instagram, Linkedin: *linkedin, Twitter: *twitter, TempatLahir: *tempat_lahir, Photo: *photo, TanggalLahir: *tanggal_lahir}
	repo.db.First(&wisudawan, "IdWisudawan = ?", id_wisudawan)
	repo.db.Model(&wisudawan).Update(wisudawan_update)
}

func (repo *WisudawanRepository) DeleteOne(id_wisudawan string) {
	var wisudawan entity.Wisudawan
	repo.db.First(&wisudawan, "IdWisudawan = ?", id_wisudawan)
	repo.db.Delete(&wisudawan)
}
