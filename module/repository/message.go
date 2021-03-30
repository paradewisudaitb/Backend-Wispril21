package repository

import (
	"gorm.io/gorm"

	"github.com/paradewisudaitb/Backend/module/entity"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) entity.MessageRepository {
	return &MessageRepository{db: db}
}
func (repo *MessageRepository) AddOne(message, sender, idWisudawan string) error {
	messageEntity := entity.Message{
		ReceiverID: idWisudawan,
		Message:    message,
		Sender:     sender,
	}
	if err := repo.db.Create(&messageEntity).Error; err != nil {
		return err
	}
	return nil
}

func (repo *MessageRepository) DeleteOne(idMessage string) error {

	if err := repo.db.First(&entity.Message{}, "id = ?", idMessage).Error; err != nil {
		return err
	}
	if err := repo.db.Where("id = ?", idMessage).Delete(&entity.Message{}).Error; err != nil {
		return err
	}
	return nil

}
func (repo *MessageRepository) GetMessage(idWisudawan string) ([]entity.Message, error) {
	var results []entity.Message
	if err := repo.db.Find(&results, "receiver_id = ?", idWisudawan).Error; err != nil {
		return make([]entity.Message, 0), err
	}
	return results, nil
}
