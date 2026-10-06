package profile

import (
	"errors"

	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IProfileService interface {
	GetProfileService(userId uint) (Profile, helpers.ErrorType, error)
	AddProfileService(model Profile) (helpers.ErrorType, error)
	UpdateProfileService(model Profile) (helpers.ErrorType, error)
	DeleteProfileService(userId uint) (helpers.ErrorType, error)
}

type ProfileService struct {
	repo IProfileRepository // Corrigido para injetar a interface do repositório
}

// AddProfileService implements [IProfileService].
func (p *ProfileService) AddProfileService(model Profile) (helpers.ErrorType, error) {
	err := p.repo.Create(model)
	if err != nil {
		return helpers.DATABASE, err
	}
	return helpers.NONE, nil
}

// DeleteProfileService implements [IProfileService].
func (p *ProfileService) DeleteProfileService(userId uint) (helpers.ErrorType, error) {
	// 1. Procura o perfil do utilizador para obter o ID primário do perfil
	profiles, err := p.repo.GetAll(userId)
	if err != nil {
		return helpers.DATABASE, err
	}
	if len(profiles) == 0 {
		return helpers.NOT_FOUND, errors.New("profile not found")
	}

	// 2. Apaga o perfil usando o ID do perfil retornado e o ID do utilizador autenticado
	err = p.repo.Delete(profiles[0].ID, userId)
	if err != nil {
		return helpers.DATABASE, err
	}
	return helpers.NONE, nil
}

// GetProfileService implements [IProfileService].
func (p *ProfileService) GetProfileService(userId uint) (Profile, helpers.ErrorType, error) {
	profiles, err := p.repo.GetAll(userId)
	if err != nil {
		return Profile{}, helpers.DATABASE, err
	}
	if len(profiles) == 0 {
		return Profile{}, helpers.NOT_FOUND, errors.New("profile not found")
	}
	return profiles[0], helpers.NONE, nil
}

// UpdateProfileService implements [IProfileService].
func (p *ProfileService) UpdateProfileService(model Profile) (helpers.ErrorType, error) {
	err := p.repo.Update(model)
	if err != nil {
		return helpers.DATABASE, err
	}
	return helpers.NONE, nil
}

// NewProfileService inicializa o serviço com a injeção do repositório.
func NewProfileService(repo IProfileRepository) IProfileService {
	return &ProfileService{
		repo: repo,
	}
}
