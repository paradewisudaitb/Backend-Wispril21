package repository

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) entity.ContentRepository {
	return ContentRepository{db: db}
}
func (repo ContentRepository) GetOne(id string) (entity.Content, error) {
	var result entity.Content
	if err := repo.db.Model(&entity.Content{}).Where("id = ?", id).First(&result).Error; err != nil {
		return entity.Content{}, err
	}
	return result, nil
}

func (repo ContentRepository) GetByWisudawan(idWisudawan string) ([]entity.Content, error) {
	var result []entity.Content
	if err := repo.db.Model(&entity.Content{}).Where("id = ?", idWisudawan).Find(&result).Error; err != nil {
		return make([]entity.Content, 0), err
	}
	return result, nil
}

func (repo ContentRepository) AddOne(idWisudawan, idOrgz, contenttype, headings, details, image string) error {
	contentEntity := entity.Content{
		WisudawanID:    idWisudawan,
		OrganizationID: idOrgz,
		Type:           contenttype,
		Headings:       headings,
		Details:        details,
		Image:          image,
	}
	if err := repo.db.Create(&contentEntity).Error; err != nil {
		return err
	}
	return nil
}

func (repo ContentRepository) UpdateOne(idContent string, idWisudawan, idOrgz, contenttype, headings, details, image string) error {
	var content entity.Content
	content_update := map[string]interface{}{}
	if idWisudawan != "" {
		content_update["wisudawan_id"] = idWisudawan
	}

	if idOrgz != "" {
		content_update["organization_id"] = idOrgz
	}

	if contenttype != "" {
		content_update["content_type"] = contenttype
	}
	if headings != "" {
		content_update["headings"] = headings
	}
	if details != "" {
		content_update["details"] = details
	}
	if image != "" {
		content_update["image"] = image
	}

	if err := repo.db.First(&entity.Wisudawan{}, "id = ?", idContent).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&content).Where("id = ?", idContent).Updates(content_update).Error; err != nil {
		return err
	}
	return nil
}

func (repo ContentRepository) DeleteOne(idContent string) error {
	if err := repo.db.First(&entity.Content{}, "id = ?", idContent).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idContent).Delete(&entity.Content{}).Error; err != nil {
		return err
	}
	return nil
}
