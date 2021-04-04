package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
)

type JurusanUseCase struct {
	jurusanrepo entity.JurusanRepository
}

func ConvertEntityJurusanToSerializer(j entity.Jurusan) entity.GetJurusanSerializer {
	return entity.GetJurusanSerializer{
		Id:            j.ID,
		Jurusan:       j.Jurusan,
		JurusanShort:  j.JurusanShort,
		Fakultas:      j.Fakultas,
		FakultasShort: j.FakultasShort,
	}
}

func NewJurusanUsecase(j entity.JurusanRepository) entity.JurusanUseCase {
	return JurusanUseCase{
		jurusanrepo: j,
	}
}

func (j JurusanUseCase) CreateJurusan(item entity.CreateJurusanSerializer) error {
	if err := j.jurusanrepo.AddOne((item.Jurusan), (item.Fakultas), (item.FakultasShort), (item.JurusanShort)); err != nil {
		return err
	}
	return nil
}

func (j JurusanUseCase) DeleteJurusan(IdJurusan uuid.UUID) error {
	err := j.jurusanrepo.DeleteOne(IdJurusan)
	if err == nil {
		return nil
	}
	return err
}

func (j JurusanUseCase) UpdateJurusan(item entity.UpdateJurusanSerializer) error {
	err := j.jurusanrepo.UpdateOne(item.IdJurusan, item.Jurusan, item.Fakultas, item.FakultasShort, item.JurusanShort)
	if err != nil {
		return err
	}
	return nil
}
func (j JurusanUseCase) GetJurusan(IdJurusan uuid.UUID) (entity.Jurusan, error) {
	result, err := j.jurusanrepo.GetOne(IdJurusan)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (j JurusanUseCase) GetAllJurusan() ([]entity.Jurusan, error) {
	result, err := j.jurusanrepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}
