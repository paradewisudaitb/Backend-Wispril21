package content

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

func (a *JurusanUseCase) CreateJurusan(item entity.CreateJurusanSerializer) error {
	a.jurusanrepo.AddOne(*item.Jurusan, *item.Fakultas, *item.FakultasShort, *item.JurusanShort)
	return nil
}

func (a *JurusanUseCase) DeleteJurusan(item entity.DeleteJurusanSerializer) error {
	err := a.jurusanrepo.DeleteOne(item.IdJurusan)
	if err == nil {
		return nil
	}
	return err
}

func (a *JurusanUseCase) UpdateJurusan(item entity.UpdateJurusanSerializer) error {
	err := a.jurusanrepo.UpdateOne(item.IdJurusan, item.Jurusan, item.Fakultas, item.FakultasShort, item.JurusanShort)
	if err == nil {
		return nil
	}
	return err
}
func (a *JurusanUseCase) GetJurusan(IdJurusan uuid.UUID) (entity.Jurusan, error) {
	jurusan, err := a.jurusanrepo.GetOne(IdJurusan)
	if err != nil {
		return jurusan, err
	}
	return jurusan, nil
}
