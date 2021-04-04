package repository

import (
	"time"

	"github.com/paradewisudaitb/Backend/module/entity"
	"gorm.io/gorm"
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

func (repo *ViewRepository) GetTop5() ([]struct {
	entity.Wisudawan
	Count int64
}, error) {
	var result []struct {
		entity.Wisudawan
		Count int64
	}
	size := repo.db.Find(&[]entity.View{}).RowsAffected
	if size > 5 {
		if err := repo.db.Raw("SELECT wisudawan.*,count from (SELECT wisudawan_id as id, count(id) as count FROM \"view\" GROUP BY \"wisudawan_id\" LIMIT 4) T INNER JOIN wisudawan ON T.id = wisudawan.id ORDER BY count desc").Scan(&result).Error; err != nil {
			return nil, err
		}
	}
	return result, nil
}
