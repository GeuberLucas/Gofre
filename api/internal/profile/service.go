package profile

import (
	"github.com/GeuberLucas/Gofre/api/pkg/helpers"
)

type IProfileService interface {
	GetProfileService(userId uint) (Profile, helpers.ErrorType, error)
	AddProfileService(model Profile) (helpers.ErrorType, error)
	UpdateProfileService(model Profile) (helpers.ErrorType, error)
	DeleteProfileService(userId uint) (helpers.ErrorType, error)
}

type ProfileService struct {
	service IProfileService
}

// AddProfileService implements [IProfileService].
func (p *ProfileService) AddProfileService(model Profile) (helpers.ErrorType, error) {
	panic("unimplemented")
}

// DeleteProfileService implements [IProfileService].
func (p *ProfileService) DeleteProfileService(userId uint) (helpers.ErrorType, error) {
	panic("unimplemented")
}

// GetProfileService implements [IProfileService].
func (p *ProfileService) GetProfileService(userId uint) (Profile, helpers.ErrorType, error) {
	panic("unimplemented")
}

// UpdateProfileService implements [IProfileService].
func (p *ProfileService) UpdateProfileService(model Profile) (helpers.ErrorType, error) {
	panic("unimplemented")
}

// AddProfileService implements [IProfileService].

func NewProfileService(service IProfileService) IProfileService {
	return &ProfileService{
		service: service,
	}
}
