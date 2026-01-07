package interfaces

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/pkg/helpers"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/models"
)

type IAggregatedRepository interface {
	InsertOrUpdate(tx *sql.Tx, model *models.Aggregated) (helpers.ErrorType, error)
	GetAll(userId int) ([]models.Aggregated, helpers.ErrorType, error)
	GetByMonthAndYear(userId int, month int, year int) (models.Aggregated, helpers.ErrorType, error)
	GetByYear(userId int, year int) ([]models.Aggregated, helpers.ErrorType, error)
}
