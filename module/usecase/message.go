package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type MessageUsecase struct {
	messagerepo entity.MessageRepository
}

func NewMessageUsecase(j entity.MessageRepository) entity.MessageUsecase {
	return MessageUsecase{
		messagerepo: j,
	}
}
func (j MessageUsecase) CreateMessage(item entity.CreateMessageSerializer) error {
	if err := j.messagerepo.AddOne((item.Message), (item.Sender), (item.IdWisudawan)); err != nil {
		return err
	}
	return nil
}

func (j MessageUsecase) DeleteMessage(idMessage uuid.UUID) error {
	if err := j.messagerepo.DeleteOne(idMessage.String()); err != nil {
		return err
	}
	return nil
}

func (j MessageUsecase) GetMessage(idWisudawan uuid.UUID) ([]entity.Message, error) {
	result, err := j.messagerepo.GetMessage(idWisudawan.String())
	if err != nil {
		return result, err
	}
	return result, nil
}
