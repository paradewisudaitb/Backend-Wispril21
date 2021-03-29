package repository

import (
	"errors"
	"reflect"

	"github.com/jinzhu/gorm"
	"github.com/paradewisudaitb/Backend/module/entity"
)

type OrgzRepository struct {
	db *gorm.DB
}

func NewOrgzRepository(db *gorm.DB) OrgzRepository {
	return OrgzRepository{db: db}
}

func (repo *OrgzRepository) GetOne(id string) (entity.Orgz, error) {
	var orgz entity.Orgz
	repo.db.First(&orgz, "id = ?", id)
	if orgz.ID == "" {
		return orgz, errors.New("Item not found")
	}
	return orgz, nil
}

func (repo *OrgzRepository) AddOne(name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string) {
	orgz := entity.Orgz{Name: name, Category: category, ApresiasiPoster: apresiasi_poster, ApresiasiTulisan: apresiasi_tulisan, ApresiasiVideo: apresiasi_video}
	repo.db.Create(&orgz)
}

// Kurang handling error buat kalau inputnya gak sesuai
func (repo *OrgzRepository) UpdateOne(id_organization string, name, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video *string) error {
	orgz := entity.Orgz{Name: *name, Category: *category, ApresiasiPoster: *apresiasi_poster, ApresiasiTulisan: *apresiasi_tulisan, ApresiasiVideo: *apresiasi_video}
	orgz_update := map[string]interface{}{}
	orgz_type := reflect.TypeOf(orgz)
	orgz_value := reflect.ValueOf(orgz)
	for i := 0; i < orgz_value.NumField(); i++ {
		field_name := orgz_type.Field(i).Name
		field_value := orgz_value.Field(i).Interface()
		if field_value == "" {
			continue
		}
		orgz_update[field_name] = field_value
	}
	var new_orgz entity.Orgz
	if orgz.ID == "" {
		return errors.New("Item not found")
	}
	repo.db.First(&new_orgz, "id = ?", id_organization)
	repo.db.Model(&new_orgz).Update(orgz_update)
	return nil
}

func (repo *OrgzRepository) DeleteOne(id_organization string) error {
	var orgz entity.Orgz
	repo.db.First(&orgz, "id = ?", id_organization)
	if orgz.ID == "" {
		return errors.New("Item not found")
	}
	repo.db.Delete(&orgz)
	return nil
}
