package statuscode

type StatusCode int

const (
	UnknownError StatusCode = iota
	Other
)

func (s StatusCode) String() string {
	return [...]string{"unknowerror", "other"}[s]
}
