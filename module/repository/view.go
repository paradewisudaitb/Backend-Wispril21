package repository

import (
	"time"

	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ViewRepository struct {
	db *gorm.DB
}

func NewViewRepository(db *gorm.DB) entity.ViewRepository {
	return &ViewRepository{db: db}
}

func (repo *ViewRepository) AddOne(idWisudawan, ip string, time time.Time) error {
	v := entity.View{
		WisudawanID: idWisudawan,
		IP:          ip,
		AccessTime:  time,
	}
	if err := repo.db.Create(&v).Error; err != nil {
		return err
	}
	return nil
}

func (repo *ViewRepository) GetLast(idWisudawan, ip string) (entity.View, error) {
	var view entity.View
	if err := repo.db.
		Order("access_time").
		Where("wisudawan_id = ?", idWisudawan).
		Where("ip = ?", ip).
		Last(&view).Error; err != nil {
		return view, err
	}
	return view, nil
}

func (repo *ViewRepository) GetTop5() ([]entity.GetViewWisudawan, error) {
	var result []entity.GetViewWisudawan
	size := repo.db.Find(&[]entity.View{}).RowsAffected
	if size > 5 {
		if err := repo.db.Raw("SELECT wisudawan.id,count from (SELECT wisudawan_id as id, count(id) as count FROM \"view\" GROUP BY \"wisudawan_id\" LIMIT 5) T INNER JOIN wisudawan ON T.id = wisudawan.id INNER JOIN jurusan ON wisudawan.jurusan_id = jurusan.id ORDER BY count desc").Scan(&result).Error; err != nil {
			return nil, err
		}
		for i := range result {
			var wisudawan entity.Wisudawan
			if err := repo.db.Preload(clause.Associations).Find(&wisudawan, "id = ?", result[i].Wisudawan.ID).Error; err != nil {
				return nil, err
			}
			result[i].Wisudawan = wisudawan
		}
	}
	return result, nil
}
