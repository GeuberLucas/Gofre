package profile

import "database/sql"

type IProfileRepository interface {
	Create(model Profile) error
	GetAll(userId uint) ([]Profile, error)
	GetById(id uint) (Profile, error)
	Update(model Profile) error
	Delete(id uint, userId uint) error
}

type ProfileRepository struct {
	db *sql.DB
}

// Create implements [IProfileRepository].
func (p *ProfileRepository) Create(model Profile) error {
	panic("unimplemented")
}

// Delete implements [IProfileRepository].
func (p *ProfileRepository) Delete(id uint, userId uint) error {
	panic("unimplemented")
}

// GetAll implements [IProfileRepository].
func (p *ProfileRepository) GetAll(userId uint) ([]Profile, error) {
	panic("unimplemented")
}

// GetById implements [IProfileRepository].
func (p *ProfileRepository) GetById(id uint) (Profile, error) {
	panic("unimplemented")
}

// Update implements [IProfileRepository].
func (p *ProfileRepository) Update(model Profile) error {
	panic("unimplemented")
}

func NewExpenseRepository(db *sql.DB) IProfileRepository {
	return &ProfileRepository{db: db}
}
