package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type ContentUsecase struct {
	contentrepo entity.ContentRepository
}

func NewContentUsecase(a entity.ContentRepository) entity.ContentUseCase {
	return ContentUsecase{
		contentrepo: a,
	}
}

func (uc ContentUsecase) CreateContent(item entity.CreateContentSerializer) error {
	if err := uc.contentrepo.AddOne(
		item.Wisudawan,
		item.Organization,
		item.ContentType,
		item.Headings,
		item.Details,
		item.Image,
	); err != nil {
		return err
	}
	return nil
}
func (uc ContentUsecase) DeleteContent(IdContent uuid.UUID) error {
	err := uc.contentrepo.DeleteOne(IdContent.String())
	if err == nil {
		return nil
	}
	return err
}

func (uc ContentUsecase) UpdateContent(item entity.UpdateContentSerializer) error {
	err := uc.contentrepo.UpdateOne(
		item.Content,
		item.Wisudawan,
		item.Organization,
		item.ContentType,
		item.Headings,
		item.Details,
		item.Image,
	)
	if err != nil {
		return err
	}
	return nil
}

func (uc ContentUsecase) GetContent(IdContent uuid.UUID) (entity.Content, error) {

	result, err := uc.contentrepo.GetOne(IdContent.String())
	if err != nil {
		return result, err
	}
	return result, nil
}

func (uc ContentUsecase) GetByWisudawan(IdWisudawan uuid.UUID) ([]entity.Content, error) {
	result, err := uc.contentrepo.GetByWisudawan(IdWisudawan.String())
	if err != nil {
		return result, err
	}
	return result, nil
}
