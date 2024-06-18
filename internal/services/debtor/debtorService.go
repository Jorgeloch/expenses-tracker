package debtorService

import (
	"errors"

	"github.com/google/uuid"
	debtorDTO "github.com/jorgeloch/expenses-tracker/internal/dto/debtor"
	debtorModel "github.com/jorgeloch/expenses-tracker/internal/models/debtor"
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

func (s *Service) GetAll(ownerID string) ([]debtorModel.Debtor, error) {
	return s.Repository.DebtorRepository.GetAll(ownerID)
}

func (s *Service) GetByID(ownerID string, id string) (debtorModel.Debtor, error) {
	return s.Repository.DebtorRepository.GetByID(ownerID, id)
}

func (s *Service) Create(ownerID string, dto debtorDTO.CreateDebtorDTO) (uuid.UUID, error) {
	location := uuid.New()

	ownerID_uuid, err := uuid.Parse(ownerID)
	if err != nil {
		return uuid.Nil, err
	}

	err = s.Repository.DebtorRepository.Create(debtorModel.Debtor{
		ID:      location,
		Name:    dto.Name,
		OwnerID: ownerID_uuid,
	})
	if err != nil {
		return uuid.Nil, err
	}
	return location, nil
}

func (s *Service) Update(ownerID string, id string, dto debtorDTO.UpdateDebtorDTO) (debtorModel.Debtor, error) {
	debtor, err := s.GetByID(ownerID, id)
	if debtor.ID == uuid.Nil {
		return debtorModel.Debtor{}, errors.New("debtor not found or debtor is not yours")
	}
	if err != nil {
		return debtorModel.Debtor{}, err
	}
	if dto.Name != "" {
		debtor.Name = dto.Name
	}
	return debtor, s.Repository.DebtorRepository.Update(debtor)
}

func (s *Service) Delete(ownerID string, id string) error {
	debtor, err := s.GetByID(ownerID, id)
	if debtor.ID == uuid.Nil {
		return errors.New("debtor not found or debtor is not yours")
	}
	if err != nil {
		return err
	}
	return s.Repository.DebtorRepository.Delete(ownerID, id)
}
