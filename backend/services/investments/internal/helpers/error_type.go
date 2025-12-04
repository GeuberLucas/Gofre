package helpers

type ErrorType int

const (
	VALIDATION ErrorType = iota
	INTERNAL
	MISSING
	STATE
	NONE
)

func (et ErrorType) String() string {
	return [...]string{"validation", "internal", "missing", "state", "none"}[et]
}
