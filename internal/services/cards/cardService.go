package cardService

import (
	"github.com/google/uuid"
	cardDTO "github.com/jorgeloch/expenses-tracker/internal/dto/card"
	cardModel "github.com/jorgeloch/expenses-tracker/internal/models/card"
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

func (s *Service) GetAll(ownerID string) ([]cardModel.Card, error) {
	return s.Repository.CardRepository.GetAll(ownerID)
}

func (s *Service) GetByID(ownerID string, id string) (cardModel.Card, error) {
	return s.Repository.CardRepository.GetByID(ownerID, id)
}

func (s *Service) Create(ownerID string, dto cardDTO.CreateCardDTO) (uuid.UUID, error) {
	location := uuid.New()
	ownerID_uuid, err := uuid.Parse(ownerID)

	err = s.Repository.CardRepository.Create(cardModel.Card{
		ID:      location,
		Name:    dto.Name,
		Flag:    dto.Flag,
		OwnerID: ownerID_uuid,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return location, nil
}

func (s *Service) Update(ownerID string, id string, dto cardDTO.UpdateCardDTO) (cardModel.Card, error) {
	card, err := s.GetByID(ownerID, id)

	if err != nil {
		return cardModel.Card{}, err
	}

	if dto.Name != "" {
		card.Name = dto.Name
	}

	if dto.Flag != "" {
		card.Flag = dto.Flag
	}

	return card, s.Repository.CardRepository.Update(card)
}

func (s *Service) Delete(ownerID string, id string) error {
	return s.Repository.CardRepository.Delete(ownerID, id)
}
