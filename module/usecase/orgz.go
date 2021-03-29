package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
)

type OrgzUseCase struct {
	orgzrepo repository.OrgzRepository
}

func NewOrgzUsecase(a repository.OrgzRepository) entity.OrgzUseCase {
	return &OrgzUseCase{
		orgzrepo: a,
	}
}

// gak tau mau handling error kalau apa
func (a *OrgzUseCase) CreateOrgz(item entity.CreateOrgzSerializer) error {
	a.orgzrepo.AddOne(item.Name, item.Category, item.Logo, *item.ApresiasiPoster, *item.ApresiasiTulisan, *item.ApresiasiVideo)
	return nil
}

func (a *OrgzUseCase) DeleteOrgz(item entity.DeleteOrgzSerializer) error {
	err := a.orgzrepo.DeleteOne(item.IdOrgz)
	if err != nil {
		return err
	}
	return nil
}

func (a *OrgzUseCase) UpdateOrgz(item entity.UpdateOrgzSerializer) error {
	err := a.orgzrepo.UpdateOne(item.IdOrgz, item.Name, item.Category, item.Logo, item.ApresiasiPoster, item.ApresiasiTulisan, item.ApresiasiVideo)
	if err != nil {
		return err
	}
	return nil
}
func (a *OrgzUseCase) GetOrgz(idOrgz string) (entity.Orgz, error) {
	orgz, err := a.orgzrepo.GetOne(idOrgz)
	if err != nil {
		return orgz, err
	}
	return orgz, nil
}
