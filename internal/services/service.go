package service

import (
	repository "github.com/jorgeloch/expenses-tracker/internal/repositories"
	"github.com/jorgeloch/expenses-tracker/internal/services/owner"
)

type Service struct {
	OwnerService *ownerService.Service
}

func Init(r *repository.Repository) *Service {
	return &Service{
		OwnerService: ownerService.Init(r),
	}
}
