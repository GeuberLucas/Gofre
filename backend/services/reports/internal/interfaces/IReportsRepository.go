package interfaces

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
)

type IReportsRepository[T any] interface {
	InsertOrUpdate(tx *sql.Tx, model *T) (helpers.ErrorType, error)
	GetAll(userId int) ([]T, helpers.ErrorType, error)
	GetByMonthAndYear(userId int, month int, year int) (T, helpers.ErrorType, error)
	GetByYear(userId int, year int) ([]T, helpers.ErrorType, error)
}
