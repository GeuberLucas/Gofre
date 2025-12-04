package types

type ErrorType int

const (
	VALIDATION ErrorType = iota
	INTERNAL
	MISSING
	STATE
)

func (et ErrorType) String() string {
	return [...]string{"validation", "internal", "missing", "state"}[et]
}
