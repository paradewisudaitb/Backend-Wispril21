package contenttype

import (
	"errors"
	"strings"
)

type ContentType int

const (
	Kontribusi ContentType = iota
	Prestasi
	Karya
	Tips
	Keanggotaan
)

func (s ContentType) String() string {
	return [...]string{"KONTRIBUSI", "PRESTASI", "KARYA", "TIPS_SUKSES", "KEANGGOTAAN"}[s]
}

func GetEnum(any string) (string, error) {
	TYPES := [...]string{"KONTRIBUSI", "PRESTASI", "KARYA", "TIPS_SUKSES", "KEANGGOTAAN"}

	anyConverted := strings.ReplaceAll(any, " ", "_")
	for i, x := range TYPES {
		if strings.EqualFold(x, anyConverted) {
			return TYPES[i], nil
		}
	}
	return "", errors.New("Unknown type")
}
