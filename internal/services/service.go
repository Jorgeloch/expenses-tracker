package service

import (
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	cardService "github.com/jorgeloch/expenses-tracker/internal/services/cards"
	debtorService "github.com/jorgeloch/expenses-tracker/internal/services/debtor"
	expenseService "github.com/jorgeloch/expenses-tracker/internal/services/expense"
	"github.com/jorgeloch/expenses-tracker/internal/services/owner"
)

type Service struct {
	OwnerService   *ownerService.Service
	CardService    *cardService.Service
	DebtorService  *debtorService.Service
	ExpenseService *expenseService.Service
}

func Init(r *repository.Repository) *Service {
	return &Service{
		OwnerService:   ownerService.Init(r),
		CardService:    cardService.Init(r),
		DebtorService:  debtorService.Init(r),
		ExpenseService: expenseService.Init(r),
	}
}
