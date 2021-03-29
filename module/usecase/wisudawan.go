package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
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
func (a WisudawanUseCase) DeleteWisudawan(item entity.DeleteWisudawanSerializer) error {
	if err := a.wisudawanrepo.DeleteOne(item.IdWisudawan); err != nil {
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

// type WisudawanRepository interface {
// 	GetOne(wisudawanID uuid.UUID) (Wisudawan, error)
// 	GetAll() ([]Wisudawan, error)
// 	AddOne(nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusan, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error
// 	UpdateOne(WisudawanID uuid.UUID, nim uint32, angkatan uint16, nama, panggilan, judulTA, jurusanID, instagram, linkedin, twitter, tempatLahir, photo string, tanggalLahir time.Time) error
// 	DeleteOne(WisudawanID uuid.UUID) error
// 	// Filter(jurusan string) ([]Wisudawan, error)
// }
func (a WisudawanUseCase) GetWisudawan(idWisudawan string) (entity.Wisudawan, error) {
	var result entity.Wisudawan
	if result, err := a.wisudawanrepo.GetOne(idWisudawan); err != nil {
		return entity.Wisudawan{}, err
	}
	return result, nil
}

func (a WisudawanUseCase) GetAllWisudawan(wisudawan string) ([]entity.Wisudawan, error) {
	var result []entity.Wisudawan
	if result, err := a.wisudawanrepo.GetOne(idWisudawan); err != nil {
		return entity.Wisudawan{}, err
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
