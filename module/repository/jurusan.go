package repository

import (
	"gorm.io/gorm"

	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type JurusanRepository struct {
	db *gorm.DB
}

func NewJurusanRepository(db *gorm.DB) entity.JurusanRepository {
	return &JurusanRepository{db: db}
}

func (repo *JurusanRepository) GetOne(id uuid.UUID) (entity.Jurusan, error) {
	var jurusan entity.Jurusan
	if err := repo.db.First(&jurusan, "id = ?", id).Error; err != nil {
		return jurusan, err
	}
	return jurusan, nil
}

func (repo *JurusanRepository) AddOne(jurusan, fakultas, fakultas_short, jurusan_short string) error {
	jurusans := entity.Jurusan{Jurusan: jurusan, Fakultas: fakultas, FakultasShort: fakultas_short, JurusanShort: jurusan_short}
	if err := repo.db.Create(&jurusans).Error; err != nil {
		return err
	}
	return nil
}

func (repo *JurusanRepository) UpdateOne(id uuid.UUID, jurusan, fakultas, fakultas_short, jurusan_short string) error {
	var target entity.Jurusan
	jurusan_update := map[string]interface{}{}
	if jurusan != "" {
		jurusan_update["jurusan"] = jurusan
	}
	if fakultas != "" {
		jurusan_update["fakultas"] = fakultas
	}
	if jurusan_short != "" {
		jurusan_update["jurusan_short"] = jurusan_short
	}
	if fakultas_short != "" {
		jurusan_update["fakultas_short"] = fakultas_short
	}
	if err := repo.db.First(&entity.Jurusan{}, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&target).Where("id = ?", id.String()).Updates(jurusan_update).Error; err != nil {
		return err
	}
	return nil
}

func (repo *JurusanRepository) DeleteOne(id uuid.UUID) error {
	if err := repo.db.First(&entity.Jurusan{}, "id = ?", id).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", id.String()).Delete(&entity.Jurusan{}).Error; err != nil {
		return err
	}
	return nil
}
