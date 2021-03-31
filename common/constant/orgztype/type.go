package contenttype

type OrgzType int

const (
	Himpunan OrgzType = iota
	Kabinet
	MWAWM
	UKM
)

func (s OrgzType) String() string {
	return [...]string{"HIMPUNAN", "KABINET", "MWAWM", "UKM"}[s]
}
