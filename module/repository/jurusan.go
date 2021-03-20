package repository

import (
	"errors"

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

func (repo *JurusanRepository) GetOne(id uuid.UUID) (entity.Jurusan, error) {
	var jurusan entity.Jurusan
	repo.db.First(&jurusan, "id = ?", id)
	if jurusan.ID == "" {
		return jurusan, errors.New("Id jurusan not found")
	}
	return jurusan, nil
}

func (repo *JurusanRepository) AddOne(jurusan, fakultas, fakultas_short, jurusan_short string) {
	jurusans := entity.Jurusan{Jurusan: jurusan, Fakultas: fakultas, FakultasShort: fakultas_short, JurusanShort: jurusan_short}
	repo.db.Create(&jurusans)
}

func (repo *JurusanRepository) UpdateOne(id_jurusan uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short *string) error {
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
	if jurusans.ID == "" {
		return errors.New("Id jurusan not found")
	}
	repo.db.First(&jurusans, "id = ?", id_jurusan)
	repo.db.Model(&jurusans).Update(jurusan_update)
	return nil
}

func (repo *JurusanRepository) DeleteOne(id_jurusan uuid.UUID) error {
	var jurusans entity.Jurusan
	repo.db.First(&jurusans, "id = ?", id_jurusan)
	if jurusans.ID == "" {
		return errors.New("Id jurusan not found")
	}
	repo.db.Delete(&jurusans)
	return nil
}
