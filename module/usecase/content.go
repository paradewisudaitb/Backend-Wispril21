package usecase

import (
	"strings"

	"github.com/paradewisudaitb/Backend/common/constant/contenttype"
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

func ConvertEntityContentsToSerializer(data []entity.Content) entity.GetContentsSerializer {
	var selfData []entity.GetContentSerializer
	var orgzData []entity.GetContentSerializer2
	for _, x := range data {
		if strings.EqualFold(x.Type, contenttype.Karya.String()) ||
			strings.EqualFold(x.Type, contenttype.Prestasi.String()) ||
			strings.EqualFold(x.Type, contenttype.Funfact.String()) ||
			strings.EqualFold(x.Type, contenttype.Tips.String()) {
			selfData = append(selfData, entity.GetContentSerializer{
				ContentType: x.Type,
				Headings:    x.Headings,
				Details:     x.Details,
				Image:       x.Image,
			})
		} else {
			orgzData = append(orgzData, entity.GetContentSerializer2{
				GetContentSerializer: entity.GetContentSerializer{
					ContentType: x.Type,
					Headings:    x.Headings,
					Details:     x.Details,
					Image:       x.Image,
				},
				OrganizationName: x.Organization.Name,
				OrganizationLogo: x.Organization.Logo,
			})
		}

	}
	if len(orgzData) == 0 {
		orgzData = make([]entity.GetContentSerializer2, 0)
	}

	if len(selfData) == 0 {
		selfData = make([]entity.GetContentSerializer, 0)
	}
	return entity.GetContentsSerializer{
		OrganizationalContents: orgzData,
		SelfContents:           selfData,
	}
}

func (uc ContentUsecase) CreateContent(item entity.CreateContentSerializer) error {
	if err := uc.contentrepo.AddOne(
		item.Nim,
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
		item.Nim,
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

func (uc ContentUsecase) GetByWisudawan(NimWisudawan uint32) ([]entity.Content, error) {
	result, err := uc.contentrepo.GetByWisudawan(NimWisudawan)
	if err != nil {
		return nil, err
	}
	return result, nil
}
