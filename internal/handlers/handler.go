package handlers

import (
	"github.com/go-playground/validator/v10"
	cardHandler "github.com/jorgeloch/expenses-tracker/internal/handlers/card"
	debtorHandler "github.com/jorgeloch/expenses-tracker/internal/handlers/debtor"
	"github.com/jorgeloch/expenses-tracker/internal/handlers/owner"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

type Handler struct {
	OwnerHandler  *ownerHandler.Handler
	CardHandler   *cardHandler.Handler
	DebtorHandler *debtorHandler.Handler
}

func Init(s *service.Service) *Handler {
	validator := validator.New()
	return &Handler{
		OwnerHandler:  ownerHandler.Init(s, validator),
		CardHandler:   cardHandler.Init(s, validator),
		DebtorHandler: debtorHandler.Init(s, validator),
	}
}
