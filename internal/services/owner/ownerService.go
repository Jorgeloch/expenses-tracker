package ownerService

import (
	"github.com/google/uuid"
	ownerDTO "github.com/jorgeloch/expenses-tracker/internal/dto/owner"
	"github.com/jorgeloch/expenses-tracker/internal/models/owner"
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	"golang.org/x/crypto/bcrypt"
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

func (s *Service) Create(dto ownerDTO.CreateOwnerDTO) (uuid.UUID, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	if err != nil {
		return uuid.Nil, err
	}

	location := uuid.New()

	err = s.Repository.OwnerRepository.Create(ownerModel.Owner{
		ID:       location,
		Email:    dto.Email,
		Name:     dto.Name,
		Password: string(password),
	})

	if err != nil {
		return uuid.Nil, err
	}

	return location, nil
}

func (s *Service) Update(user ownerModel.Owner) error {
	return s.Repository.OwnerRepository.Update(user)
}

func (s *Service) Delete(id int) error {
	return s.Repository.OwnerRepository.Delete(id)
}
