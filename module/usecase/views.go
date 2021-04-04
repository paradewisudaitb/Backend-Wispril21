// // entity
// // entitybase, id_wisudawan, ip, time

// // Usecase AddView

// // Repository AddView

// // request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record

package usecase

import (
	"github.com/paradewisudaitb/Backend/module/entity"
)

type ViewUseCase struct {
	Viewrepo entity.ViewRepository
}

// func NewViewUsecase(v entity.ViewRepository) entity.ViewUseCase {
// 	return ViewUseCase{
// 		Viewrepo: v,
// 	}
// }

// func (uc ViewUseCase) AddView(item entity.ViewSerializer) error {
// 	if err := uc.Viewrepo.AddOne(
// 		item.IdWisudawan,
// 		item.IP,
// 		item.Time,
// 	); err != nil {
// 		return err
// 	}
// 	return nil
// }
