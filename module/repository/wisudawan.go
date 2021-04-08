package repository

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/paradewisudaitb/Backend/module/entity"
)

type WisudawanRepository struct {
	db *gorm.DB
}

func NewWisudawanRepository(db *gorm.DB) WisudawanRepository {
	return WisudawanRepository{db: db}
}
func (repo WisudawanRepository) GetOne(wisudawanID string) (entity.Wisudawan, error) {
	var result entity.Wisudawan
	if err := repo.db.Preload(clause.Associations).Where("id = ?", wisudawanID).First(&result).Error; err != nil {
		return entity.Wisudawan{}, err
	}
	return result, nil
}

func (repo WisudawanRepository) GetAll() ([]entity.Wisudawan, error) {
	var allResults []entity.Wisudawan

	if err := repo.db.Find(&allResults).Error; err != nil {
		return make([]entity.Wisudawan, 0), err
	}
	return allResults, nil
}
func (repo WisudawanRepository) AddOne(nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusanID, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error {
	wisudawan := entity.Wisudawan{
		Nim:         nim,
		Angkatan:    angkatan,
		Nama:        nama,
		Panggilan:   panggilan,
		JudulTA:     judulTA,
		Instagram:   instagram,
		Linkedin:    linkedin,
		Twitter:     twitter,
		TempatLahir: tempatLahir,
		Photo:       photo,
		JurusanID:   jurusanID,
	}

	if !tanggalLahir.IsZero() {
		wisudawan.TanggalLahir = tanggalLahir
	}
	if err := repo.db.Create(&wisudawan).Error; err != nil {
		return err
	}
	return nil
}

func (repo WisudawanRepository) UpdateOne(nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusanID, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error {
	var wisudawan entity.Wisudawan
	wisudawan_update := map[string]interface{}{}
	if angkatan != 0 {
		wisudawan_update["angkatan"] = angkatan
	}
	if nama != "" {
		wisudawan_update["nama"] = nama
	}
	if panggilan != "" {
		wisudawan_update["panggilan"] = panggilan
	}
	if judulTA != "" {
		wisudawan_update["judul_ta"] = judulTA
	}
	if instagram != "" {
		wisudawan_update["instagram"] = instagram
	}
	if linkedin != "" {
		wisudawan_update["linkedin"] = linkedin
	}
	if twitter != "" {
		wisudawan_update["twitter"] = twitter
	}
	if tempatLahir != "" {
		wisudawan_update["tempat_lahir"] = tempatLahir
	}
	if photo != "" {
		wisudawan_update["photo"] = photo
	}
	if !tanggalLahir.IsZero() {
		wisudawan_update["tanggal_lahir"] = tanggalLahir
	}
	if jurusanID != "" {
		wisudawan_update["jurusan_id"] = jurusanID
	}
	if err := repo.db.First(&entity.Wisudawan{}, "nim = ?", nim).Error; err != nil {
		return err
	}
	if err := repo.db.Model(&wisudawan).Where("nim = ?", nim).Updates(wisudawan_update).Error; err != nil {
		return err
	}
	return nil
}

func (repo WisudawanRepository) DeleteOne(wisudawanID string) error {
	if err := repo.db.First(&entity.Wisudawan{}, "id = ?", wisudawanID).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", wisudawanID).Delete(&entity.Wisudawan{}).Error; err != nil {
		return err
	}
	return nil
}

func (repo WisudawanRepository) FilterByOrgzSlug(organizationSlug string) ([]entity.Wisudawan, error) {
	var result []entity.Wisudawan

	resultWisudawanID := repo.db.Table("wisudawan").
		Joins("INNER JOIN content ON content.wisudawan_id = wisudawan.id").
		Joins("INNER JOIN organization ON organization.id = content.organization_id").
		Where("slug = ?", organizationSlug).
		Distinct("wisudawan.id")

	if err := resultWisudawanID.Error; err != nil {
		return nil, err
	}

	if err := repo.db.Preload(clause.Associations).
		Find(&result, "id IN (?)", resultWisudawanID).
		Error; err != nil {
		return nil, err
	}
	return result, nil
}
