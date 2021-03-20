package content

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
)

type WisudawanUseCase struct {
	wisudawanrepo repository.WisudawanRepository
}

func NewWisudawanUsecase(a repository.WisudawanRepository) entity.WisudawanUsecase {
	return &WisudawanUseCase{
		wisudawanrepo: a,
	}
}

func (a *WisudawanUseCase) CreateWisudawan(item entity.CreateWisudawanSerializer) error
func (a *WisudawanUseCase) DeleteWisudawan(item entity.DeleteWisudawanSerializer) error
func (a *WisudawanUseCase) UpdateWisudawan(item entity.UpdateWisudawanSerializer) error
func (a *WisudawanUseCase) GetWisudawan(idWisudawan string) error
func (a *WisudawanUseCase) GetAllWisudawan(Wisudawan string) error //
func (a *WisudawanUseCase) FilterWisudawan(jurusan string) ([]entity.Wisudawan, error) {
	wisudawan := a.wisudawanrepo.Filter(jurusan)
	return wisudawan, nil
	// ini juga gak tau errornya buat apa
}
