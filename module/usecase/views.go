// // entity
// // entitybase, id_wisudawan, ip, time

// // Usecase AddView

// // Repository AddView

// // request masuk -> ambil ip address sama id_wisudawan -> dicek apakah record sudah ada di tabel -> tambah record

package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/paradewisudaitb/Backend/module/entity"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type ViewUseCase struct {
	Viewrepo entity.ViewRepository
}

func NewViewUsecase(v entity.ViewRepository) entity.ViewUseCase {
	return ViewUseCase{
		Viewrepo: v,
	}
}

func (uc ViewUseCase) AddView(idWisudawan uuid.UUID, clientIP string) error {
	lastRecord, lastErr := uc.Viewrepo.GetLast(idWisudawan.String(), clientIP)
	if lastErr != nil {
		if !errors.Is(lastErr, gorm.ErrRecordNotFound) {
			return lastErr
		}
	} else {
		diff := time.Now().Sub(lastRecord.AccessTime).Minutes()
		fmt.Println(diff)
		if diff < 10 {
			return nil
		}
	}

	if err := uc.Viewrepo.AddOne(
		idWisudawan.String(),
		clientIP,
		time.Now(),
	); err != nil {
		return err
	}
	return nil
}

func (uc ViewUseCase) GetTop5() ([]entity.GetViewWisudawan, error) {
	result, err := uc.Viewrepo.GetTop5()
	if err != nil {
		return result, err
	}
	return result, nil
}
