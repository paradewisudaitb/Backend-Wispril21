package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
	"github.com/paradewisudaitb/Backend/module/repository"
	uuid "github.com/satori/go.uuid"
)

type JurusanUseCase struct {
	jurusanrepo repository.JurusanRepository
}

func NewJurusanUsecase(j repository.JurusanRepository) entity.JurusanUseCase {
	return &JurusanUseCase{
		jurusanrepo: j,
	}
}

func (j *JurusanUseCase) CreateJurusan(item entity.CreateJurusanSerializer) error {
	j.jurusanrepo.AddOne((item.Jurusan), (item.Fakultas), (item.FakultasShort), (item.JurusanShort))
	return nil
}

func (j *JurusanUseCase) DeleteJurusan(IdJurusan uuid.UUID) error {
	err := j.jurusanrepo.DeleteOne(IdJurusan)
	if err == nil {
		return nil
	}
	return err
}

func (j *JurusanUseCase) UpdateJurusan(item entity.UpdateJurusanSerializer) error {
	err := j.jurusanrepo.UpdateOne(item.IdJurusan, item.Jurusan, item.Fakultas, item.FakultasShort, item.JurusanShort)
	if err != nil {
		return err
	}
	return nil
}
func (j *JurusanUseCase) GetJurusan(IdJurusan uuid.UUID) (entity.Jurusan, error) {
	jurusans, err := j.jurusanrepo.GetOne(IdJurusan)
	if err != nil {
		return jurusans, err
	}
	return jurusans, nil
}
