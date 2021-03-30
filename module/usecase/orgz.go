package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type OrgzUseCase struct {
	orgzrepo entity.OrgzRepository
}

func NewOrgzUsecase(a entity.OrgzRepository) entity.OrgzUseCase {
	return OrgzUseCase{
		orgzrepo: a,
	}
}

func (uc OrgzUseCase) CreateOrgz(item entity.CreateOrgzSerializer) error {
	if err := uc.orgzrepo.AddOne(
		item.Name,
		item.Slug,
		item.Category,
		item.Logo,
		item.ApresiasiPoster,
		item.ApresiasiTulisan,
		item.ApresiasiVideo); err != nil {
		return err
	}
	return nil
}

func (uc OrgzUseCase) DeleteOrgz(idOrgz uuid.UUID) error {
	err := uc.orgzrepo.DeleteOne(idOrgz.String())
	if err == nil {
		return nil
	}
	return err
}

func (uc OrgzUseCase) UpdateOrgz(item entity.UpdateOrgzSerializer) error {
	err := uc.orgzrepo.UpdateOne(
		item.IdOrgz,
		item.Name,
		item.Slug,
		item.Category,
		item.Logo,
		item.ApresiasiPoster,
		item.ApresiasiTulisan,
		item.ApresiasiVideo,
	)
	if err != nil {
		return err
	}
	return nil
}

func (uc OrgzUseCase) GetOrgz(idOrgz uuid.UUID) (entity.Orgz, error) {
	result, err := uc.orgzrepo.GetOne(idOrgz.String())
	if err != nil {
		return result, err
	}
	return result, nil
}
