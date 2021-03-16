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

func (repo *WisudawanRepository) AddOne(nim uint32, angkatan uint16, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo string, tanggal_lahir time.Time) {
	wisudawan := entity.Wisudawan{Nim: nim, Angkatan: angkatan, Nama: nama, Panggilan: panggilan, JudulTA: judul_ta, Instagram: instagram, Linkedin: linkedin, Twitter: twitter, TempatLahir: tempat_lahir, Photo: photo, TanggalLahir: tanggal_lahir}
	j := entity.Jurusan{}
	repo.db.First(&j, "id = ?", jurusan)
	wisudawan.Jurusan = j
	repo.db.Create(&wisudawan)
}

func (repo *WisudawanRepository) UpdateOne(id_wisudawan string, nim *uint32, angkatan *uint16, nama, panggilan, judul_ta, jurusan, instagram, linkedin, twitter, tempat_lahir, photo *string, tanggal_lahir *time.Time) {
	var wisudawan entity.Wisudawan
	wisudawan_update := map[string]interface{}{}
	if nim != nil {
		wisudawan_update["nim"] = *nim
	}
	if angkatan != nil {
		wisudawan_update["angkatan"] = *angkatan
	}
	if nama != nil {
		wisudawan_update["nama"] = *nama
	}
	if panggilan != nil {
		wisudawan_update["panggilan"] = *panggilan
	}
	if judul_ta != nil {
		wisudawan_update["judul_ta"] = *judul_ta
	}
	if jurusan != nil {
		wisudawan_update["jurusan"] = *jurusan
	}
	if instagram != nil {
		wisudawan_update["instagram"] = *instagram
	}
	if linkedin != nil {
		wisudawan_update["linkedin"] = *linkedin
	}
	if twitter != nil {
		wisudawan_update["twitter"] = *twitter
	}
	if tempat_lahir != nil {
		wisudawan_update["tempat_lahir"] = *tempat_lahir
	}
	if photo != nil {
		wisudawan_update["photo"] = *photo
	}
	if tanggal_lahir != nil {
		wisudawan_update["tanggal_lahir"] = *tanggal_lahir
	}
	repo.db.First(&wisudawan, "id = ?", id_wisudawan)
	repo.db.Model(&wisudawan).Update(wisudawan_update)
}

func (repo *WisudawanRepository) DeleteOne(id_wisudawan string) {
	var wisudawan entity.Wisudawan
	repo.db.First(&wisudawan, "IdWisudawan = ?", id_wisudawan)
	repo.db.Delete(&wisudawan)
}
