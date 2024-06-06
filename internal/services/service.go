package service

import (
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	cardService "github.com/jorgeloch/expenses-tracker/internal/services/cards"
	debtorService "github.com/jorgeloch/expenses-tracker/internal/services/debtor"
	"github.com/jorgeloch/expenses-tracker/internal/services/owner"
)

type Service struct {
	OwnerService  *ownerService.Service
	CardService   *cardService.Service
	DebtorService *debtorService.Service
}

func Init(r *repository.Repository) *Service {
	return &Service{
		OwnerService:  ownerService.Init(r),
		CardService:   cardService.Init(r),
		DebtorService: debtorService.Init(r),
	}
}
