package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type WisudawanUseCase struct {
	wisudawanrepo entity.WisudawanRepository
}

func NewWisudawanUsecase(a entity.WisudawanRepository) entity.WisudawanUsecase {
	return WisudawanUseCase{
		wisudawanrepo: a,
	}

}

func (a WisudawanUseCase) CreateWisudawan(item entity.CreateWisudawanSerializer) error {
	if err := a.wisudawanrepo.AddOne(
		item.Nim,
		item.Angkatan,
		item.Nama,
		item.Panggilan,
		item.JudulTA,
		item.Jurusan,
		item.Instagram,
		item.Linkedin,
		item.Twitter,
		item.TempatLahir,
		item.Photo,
		item.TanggalLahir,
	); err != nil {
		return err
	}
	return nil
}
func (a WisudawanUseCase) DeleteWisudawan(idWisudawan uuid.UUID) error {
	if err := a.wisudawanrepo.DeleteOne(idWisudawan); err != nil {
		return err
	}
	return nil
}

func (a WisudawanUseCase) UpdateWisudawan(item entity.UpdateWisudawanSerializer) error {
	if err := a.wisudawanrepo.UpdateOne(
		item.IdWisudawan,
		item.NIM,
		item.Angkatan,
		item.Nama,
		item.Panggilan,
		item.JudulTA,
		item.Jurusan,
		item.Instagram,
		item.Linkedin,
		item.Twitter,
		item.TempatLahir,
		item.Photo,
		item.TanggalLahir,
	); err != nil {
		return err
	}
	return nil
}

func (a WisudawanUseCase) GetWisudawan(idWisudawan uuid.UUID) (entity.Wisudawan, error) {
	var result entity.Wisudawan
	if result, err := a.wisudawanrepo.GetOne(idWisudawan); err != nil {
		return result, err
	}
	return result, nil
}

func (a WisudawanUseCase) GetAllWisudawan() ([]entity.Wisudawan, error) {
	var result []entity.Wisudawan
	if result, err := a.wisudawanrepo.GetAll(); err != nil {
		return result, err
	}
	return result, nil
}
func (a WisudawanUseCase) FilterWisudawan(jurusan string) ([]entity.Wisudawan, error) {
	var temp []entity.Wisudawan

	return temp, nil
}

// func (a *WisudawanUseCase) CreateWisudawan(item entity.CreateWisudawanSerializer) error
// func (a *WisudawanUseCase) DeleteWisudawan(item entity.DeleteWisudawanSerializer) error
// func (a *WisudawanUseCase) UpdateWisudawan(item entity.UpdateWisudawanSerializer) error
// func (a *WisudawanUseCase) GetWisudawan(idWisudawan string) error
// func (a *WisudawanUseCase) GetAllWisudawan(Wisudawan string) error //
// func (a *WisudawanUseCase) FilterWisudawan(jurusan string) ([]entity.Wisudawan, error) {
// 	wisudawan := a.wisudawanrepo.Filter(jurusan)
// 	return wisudawan, nil
// 	// ini juga gak tau errornya buat apa
// }
