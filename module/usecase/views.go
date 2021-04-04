// // entity
// // entitybase, id_wisudawan, ip, time

// // Usecase AddViews

// // Repository AddViews

// // request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record

package usecase

// import (
// 	"github.com/paradewisudaitb/Backend/module/entity"
// )

// type ViewsUseCase struct {
// 	viewsrepo entity.ViewsRepository
// }

// func NewViewsUsecase(v entity.ViewsRepository) entity.ViewsUseCase {
// 	return ViewsUseCase{
// 		viewsrepo: v,
// 	}
// }

// func (uc ViewsUseCase) AddViews(item entity.ViewsSerializer) error {
// 	if err := uc.viewsrepo.AddOne(
// 		item.IdWisudawan,
// 		item.IP,
// 		item.Time,
// 	); err != nil {
// 		return err
// 	}
// 	return nil
// }
