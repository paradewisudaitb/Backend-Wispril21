package repository

// import (
// 	"time"

// 	"github.com/paradewisudaitb/Backend/module/entity"
// 	uuid "github.com/satori/go.uuid"
// 	"gorm.io/gorm"
// )

// type ViewsRepository struct {
// 	db *gorm.DB
// }

// func NewViewsRepository(db *gorm.DB) entity.ViewsRepository {
// 	return &ViewsRepository{db: db}
// }

// func (repo *ViewsRepository) AddOne(idWisudawan uuid.UUID, ip string, time time.Time) error {
// 	var views entity.Views
// 	if err := repo.db.Model(&entity.Views{}).Where("IP = ?", ip).FirstOrInit(&views).Error; err != nil {
// 		return err
// 	}
// 	viewsEntity := entity.Views{
// 		IdWisudawan: idWisudawan,
// 		IP:          ip,
// 		Time:        time,
// 	}
// 	if err := repo.db.Create(&viewsEntity).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
