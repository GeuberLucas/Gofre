package repository

import (
	"database/sql"

	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/models"
)

type IPortfolioRepository interface {
	Create(model models.Portfolio) error
	GetAll() ([]models.Portfolio, error)
	GetById(id int64) (models.Portfolio, error)
	GetByUserId(userId int64) ([]models.Portfolio, error)
	Update(model models.Portfolio) error
	Delete(id int64, userId int64) error
}

type PortfolioRepository struct {
	db *sql.DB
}

// Create implements IPortfolioRepository.
func (p *PortfolioRepository) Create(model models.Portfolio) error {
	panic("unimplemented")
}

// Delete implements IPortfolioRepository.
func (p *PortfolioRepository) Delete(id int64, userId int64) error {
	panic("unimplemented")
}

// GetAll implements IPortfolioRepository.
func (p *PortfolioRepository) GetAll() ([]models.Portfolio, error) {
	panic("unimplemented")
}

// GetById implements IPortfolioRepository.
func (p *PortfolioRepository) GetById(id int64) (models.Portfolio, error) {
	panic("unimplemented")
}

// GetByUserId implements IPortfolioRepository.
func (p *PortfolioRepository) GetByUserId(userId int64) ([]models.Portfolio, error) {
	panic("unimplemented")
}

// Update implements IPortfolioRepository.
func (p *PortfolioRepository) Update(model models.Portfolio) error {
	panic("unimplemented")
}

func NewPortfolioRepository(db *sql.DB) IPortfolioRepository {
	return &PortfolioRepository{db: db}
}
