package repository

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (repo ContentRepository) GetByWisudawan(nimWisudawan uint32) ([]entity.Content, error) {
	var wisudawan entity.Wisudawan
	if err := repo.db.Model(&entity.Wisudawan{}).Find(&wisudawan, "nim = ?", nimWisudawan).Error; err != nil {
		return nil, err
	}

	var result []entity.Content
	if err := repo.db.Preload(clause.Associations).
		Where("wisudawan_id = ?", wisudawan.ID).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo ContentRepository) AddOne(nimWisudawan uint32, idOrgz, contenttype, headings, details, image string) error {
	var wisudawan entity.Wisudawan
	if err := repo.db.Model(&entity.Wisudawan{}).Where("nim = ?", nimWisudawan).First(&wisudawan).Error; err != nil {
		return err
	}

	contentEntity := entity.Content{
		WisudawanID:    wisudawan.ID,
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

func (repo ContentRepository) UpdateOne(idContent string, nimWisudawan uint32, idOrgz, contenttype, headings, details, image string) error {
	var content entity.Content
	content_update := map[string]interface{}{}
	if nimWisudawan != 0 {
		var wisudawan entity.Wisudawan
		if err := repo.db.Model(&entity.Wisudawan{}).Find(&wisudawan, "nim = ?", nimWisudawan).Error; err != nil {
			return err
		}
		content_update["wisudawan_id"] = wisudawan.ID
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
