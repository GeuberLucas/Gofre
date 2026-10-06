package profile

import (
	"database/sql"
)

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
	query := `
		INSERT INTO profiles.user_profiles (user_id, full_name, mobile_phone, initial_balance) 
		VALUES ($1, $2, $3, $4) RETURNING id`
	err := p.db.QueryRow(query, model.UserId, model.CompleteName, model.CellPhone, model.InitialBalance).Scan(&model.ID)
	return err
}

// Delete implements [IProfileRepository].
func (p *ProfileRepository) Delete(id uint, userId uint) error {
	query := `DELETE FROM profiles.user_profiles WHERE id = $1 AND user_id = $2`
	_, err := p.db.Exec(query, id, userId)
	return err
}

// GetAll implements [IProfileRepository].
// Como o user_id é UNIQUE na base de dados, isto retornará no máximo um perfil.
func (p *ProfileRepository) GetAll(userId uint) ([]Profile, error) {
	var profiles []Profile
	query := `SELECT id, user_id, full_name, mobile_phone, initial_balance FROM profiles.user_profiles WHERE user_id = $1`
	rows, err := p.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var profile Profile
		if err := rows.Scan(&profile.ID, &profile.UserId, &profile.CompleteName, &profile.CellPhone, &profile.InitialBalance); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	return profiles, nil
}

// GetById implements [IProfileRepository].
func (p *ProfileRepository) GetById(id uint) (Profile, error) {
	var profile Profile
	query := `SELECT id, user_id, full_name, mobile_phone, initial_balance FROM profiles.user_profiles WHERE id = $1`
	err := p.db.QueryRow(query, id).Scan(&profile.ID, &profile.UserId, &profile.CompleteName, &profile.CellPhone, &profile.InitialBalance)
	return profile, err
}

// Update implements [IProfileRepository].
func (p *ProfileRepository) Update(model Profile) error {
	query := `
		UPDATE profiles.user_profiles 
		SET full_name = $1, mobile_phone = $2, initial_balance = $3 
		WHERE id = $4 AND user_id = $5`
	_, err := p.db.Exec(query, model.CompleteName, model.CellPhone, model.InitialBalance, model.ID, model.UserId)
	return err
}

// Corrigido de NewExpenseRepository para NewProfileRepository
func NewProfileRepository(db *sql.DB) IProfileRepository {
	return &ProfileRepository{db: db}
}
