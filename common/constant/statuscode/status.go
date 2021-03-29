package statuscode

type StatusCode int

const (
	UnknownError StatusCode = iota
	UncompatibleJSON
	OK
	EmptyParam
	UnknownUUID
)

func (s StatusCode) String() string {
	return [...]string{"unknowerror", "uncompatible_json", "ok", "empty_param", "unknown_uuid"}[s]
}
