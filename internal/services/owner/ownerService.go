package ownerService

import (
	"github.com/google/uuid"
	"github.com/jorgeloch/expenses-tracker/cmd/api/models/owner"
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
)

type Service struct {
	Repository *repository.Repository
}

func Init(r *repository.Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) GetAll() ([]ownerModel.Owner, error) {
	return s.Repository.OwnerRepository.GetAll()
}

func (s *Service) GetByID(id int) (ownerModel.Owner, error) {
	return s.Repository.OwnerRepository.GetByID(id)
}

func (s *Service) Create(user ownerModel.Owner) (uuid.UUID, error) {
	return s.Repository.OwnerRepository.Create(user)
}

func (s *Service) Update(user ownerModel.Owner) error {
	return s.Repository.OwnerRepository.Update(user)
}

func (s *Service) Delete(id int) error {
	return s.Repository.OwnerRepository.Delete(id)
}
