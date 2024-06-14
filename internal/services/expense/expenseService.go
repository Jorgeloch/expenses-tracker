package expenseService

import (
	"github.com/google/uuid"
	expenseDTO "github.com/jorgeloch/expenses-tracker/internal/dto/expense"
	expenseModel "github.com/jorgeloch/expenses-tracker/internal/models/expense"
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

func (s *Service) GetAll(ownerID string) ([]expenseModel.Expense, error) {
	return s.Repository.ExpenseRepository.GetAll(ownerID)
}

func (s *Service) GetByID(ownerID string, id string) (expenseModel.Expense, error) {
	return s.Repository.ExpenseRepository.GetByID(ownerID, id)
}

func (s *Service) Create(ownerID string, dto expenseDTO.CreateExpenseDTO) (uuid.UUID, error) {
	location := uuid.New()
	ownerID_uuid, err := uuid.Parse(ownerID)

	err = s.Repository.ExpenseRepository.Create(expenseModel.Expense{
		ID:          location,
		Value:       dto.Value,
		Description: dto.Description,
		Date:        dto.Date,
		OwnerID:     ownerID_uuid,
		DebtorID:    dto.DebtorID,
		CardID:      dto.CardID,
	})

	if err != nil {
		return uuid.Nil, err
	}

	return location, nil
}

func (s *Service) Update(ownerID string, id string, dto expenseDTO.UpdateExpenseDTO) (expenseModel.Expense, error) {
	expense, err := s.GetByID(ownerID, id)

	if err != nil {
		return expenseModel.Expense{}, err
	}

	if dto.Value != 0 {
		expense.Value = dto.Value
	}

	if dto.Description != "" {
		expense.Description = dto.Description
	}

	if !dto.Date.IsZero() {
		expense.Date = dto.Date
	}

	if dto.DebtorID != uuid.Nil {
		expense.DebtorID = dto.DebtorID
	}

	if dto.CardID != uuid.Nil {
		expense.CardID = dto.CardID
	}

	return expense, s.Repository.ExpenseRepository.Update(expense)
}

func (s *Service) Delete(ownerID string, id string) error {
	return s.Repository.ExpenseRepository.Delete(ownerID, id)
}
