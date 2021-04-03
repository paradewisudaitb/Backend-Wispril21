package usecase

import (
	"time"

	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type WisudawanUseCase struct {
	wisudawanrepo entity.WisudawanRepository
}

func ConvertEntityWisudawanToSerializer(x entity.Wisudawan) entity.GetWisudawanSerializer {
	return entity.GetWisudawanSerializer{
		ID:            x.ID,
		Nim:           x.Nim,
		Nama:          x.Nama,
		Panggilan:     x.Panggilan,
		JudulTA:       x.JudulTA,
		Angkatan:      x.Angkatan,
		Jurusan:       x.Jurusan.Jurusan,
		JurusanShort:  x.Jurusan.JurusanShort,
		Fakultas:      x.Jurusan.Fakultas,
		FakultasShort: x.Jurusan.FakultasShort,
		Instagram:     x.Instagram,
		Linkedin:      x.Linkedin,
		Twitter:       x.Twitter,
		TempatLahir:   x.TempatLahir,
		TanggalLahir:  x.TanggalLahir.Format("02-01-2006"),
		Photo:         x.Photo,
	}
}
func ConvertEntityWisudawanToSimpleSerializer(x entity.Wisudawan) entity.GetSimpleWisudawanSerializer {
	return entity.GetSimpleWisudawanSerializer{
		ID:            x.ID,
		Nim:           x.Nim,
		Nama:          x.Nama,
		JudulTA:       x.JudulTA,
		Jurusan:       x.Jurusan.Jurusan,
		JurusanShort:  x.Jurusan.JurusanShort,
		Fakultas:      x.Jurusan.Fakultas,
		FakultasShort: x.Jurusan.FakultasShort,
		Photo:         x.Photo,
	}
}

func NewWisudawanUsecase(a entity.WisudawanRepository) entity.WisudawanUsecase {
	return WisudawanUseCase{
		wisudawanrepo: a,
	}

}

func (a WisudawanUseCase) CreateWisudawan(item entity.CreateWisudawanSerializer) error {
	tglLahir, timeErr := time.Parse("01-02-2006", item.TanggalLahir)
	if timeErr != nil {
		tglLahir = time.Time{}
	}
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
		tglLahir,
	); err != nil {
		return err
	}
	return nil
}
func (a WisudawanUseCase) DeleteWisudawan(idWisudawan uuid.UUID) error {
	if err := a.wisudawanrepo.DeleteOne(idWisudawan.String()); err != nil {
		return err
	}
	return nil
}

func (a WisudawanUseCase) UpdateWisudawan(item entity.UpdateWisudawanSerializer) error {
	tglLahir, timeErr := time.Parse("01-02-2006", item.TanggalLahir)
	if timeErr != nil {
		tglLahir = time.Time{}
	}
	if err := a.wisudawanrepo.UpdateOne(
		item.IdWisudawan.String(),
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
		tglLahir,
	); err != nil {
		return err
	}
	return nil
}

func (a WisudawanUseCase) GetWisudawan(idWisudawan uuid.UUID) (entity.Wisudawan, error) {
	result, err := a.wisudawanrepo.GetOne(idWisudawan.String())
	if err != nil {
		return result, err
	}
	return result, nil
}

func (a WisudawanUseCase) GetAllWisudawan() ([]entity.Wisudawan, error) {
	result, err := a.wisudawanrepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}
func (a WisudawanUseCase) FilterWisudawanByOrgzSlug(organizationSlug string) ([]entity.Wisudawan, error) {
	result, err := a.wisudawanrepo.FilterByOrgzSlug(organizationSlug)
	if err != nil {
		return result, err
	}
	return result, nil
}
