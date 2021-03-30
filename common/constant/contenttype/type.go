package contenttype

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
