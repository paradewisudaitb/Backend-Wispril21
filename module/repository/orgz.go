package repository

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
)

type OrgzRepository struct {
	db *gorm.DB
}

func NewOrgzRepository(db *gorm.DB) OrgzRepository {
	return OrgzRepository{db: db}
}

func (repo OrgzRepository) GetOne(idOrgz string) (entity.Orgz, error) {
	var result entity.Orgz
	if err := repo.db.First(&result, "id = ?", idOrgz).Error; err != nil {
		return result, err
	}
	return result, nil
}

func (repo OrgzRepository) AddOne(name, slug, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string) error {
	create := entity.Orgz{
		Name:             name,
		Slug:             slug,
		Category:         category,
		Logo:             logo,
		ApresiasiPoster:  apresiasi_poster,
		ApresiasiTulisan: apresiasi_tulisan,
		ApresiasiVideo:   apresiasi_video,
	}
	if err := repo.db.Create(&create).Error; err != nil {
		return err
	}
	return nil
}

func (repo OrgzRepository) UpdateOne(idOrgz, name, slug, category, logo, apresiasi_poster, apresiasi_tulisan, apresiasi_video string) error {
	var target entity.Jurusan
	update := map[string]interface{}{}
	if idOrgz != "" {
		update["id"] = idOrgz
	}
	if slug != "" {
		update["slug"] = slug
	}
	if name != "" {
		update["name"] = name
	}
	if category != "" {
		update["category"] = category
	}
	if logo != "" {
		update["logo"] = logo
	}
	if apresiasi_poster != "" {
		update["apresiasi_poster"] = apresiasi_poster
	}
	if apresiasi_tulisan != "" {
		update["apresiasi_tulisan"] = apresiasi_tulisan
	}
	if apresiasi_video != "" {
		update["apresiasi_video"] = apresiasi_video
	}
	if err := repo.db.First(&entity.Orgz{}, "id = ?", idOrgz).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&target).Where("id = ?", idOrgz).Updates(update).Error; err != nil {
		return err
	}
	return nil
}

func (repo OrgzRepository) DeleteOne(idOrgz string) error {
	if err := repo.db.First(&entity.Orgz{}, "id = ?", idOrgz).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idOrgz).Delete(&entity.Orgz{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo OrgzRepository) GetAll() ([]entity.Orgz, error) {
	var results []entity.Orgz
	if err := repo.db.Find(&results).Error; err != nil {
		return make([]entity.Orgz, 0), err
	}
	return results, nil
}
func (repo OrgzRepository) GetBySlug(slug string) (entity.Orgz, error) {
	var result entity.Orgz
	if err := repo.db.First(&result, "slug = ?", slug).Error; err != nil {
		return entity.Orgz{}, err
	}
	return result, nil
}
