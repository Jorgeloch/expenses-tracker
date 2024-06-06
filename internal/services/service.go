package service

import (
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	cardService "github.com/jorgeloch/expenses-tracker/internal/services/cards"
	"github.com/jorgeloch/expenses-tracker/internal/services/owner"
)

type Service struct {
	OwnerService *ownerService.Service
	CardService  *cardService.Service
}

func Init(r *repository.Repository) *Service {
	return &Service{
		OwnerService: ownerService.Init(r),
		CardService:  cardService.Init(r),
	}
}
