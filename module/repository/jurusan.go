package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type JurusanRepository struct {
	db *gorm.DB
}

func NewJurusanRepository(db *gorm.DB) JurusanRepository {
	return JurusanRepository{db: db}
}

func (repo *JurusanRepository) AddOne(id_jurusan uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short string) {
	jurusans := entity.Jurusan{IdJurusan: id_jurusan, Jurusan: jurusan, Fakultas: fakultas, FakultasShort: fakultas_short, JurusanShort: jurusan_short}
	repo.db.Create(&jurusans)
}

func (repo *JurusanRepository) UpdateOne(id_jurusan uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short *string) {
	var jurusans entity.Jurusan
	jurusan_update := map[string]interface{}{}
	if jurusan != nil {
		jurusan_update["jurusan"] = *jurusan
	}
	if fakultas != nil {
		jurusan_update["fakultas"] = *fakultas
	}
	if jurusan_short != nil {
		jurusan_update["jurusan_short"] = *jurusan_short
	}
	if fakultas_short != nil {
		jurusan_update["fakultas_short"] = *fakultas_short
	}
	repo.db.First(&jurusans, "id = ?", id_jurusan)
	repo.db.Model(&jurusans).Update(jurusan_update)
}

func (repo *JurusanRepository) DeleteOne(id_jurusan uuid.UUID) {
	var jurusans entity.Jurusan
	repo.db.First(&jurusans, "IdJurusan = ?", id_jurusan)
	repo.db.Delete(&jurusans)
}
