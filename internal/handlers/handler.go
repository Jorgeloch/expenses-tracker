package handlers

import (
	"log"

	"github.com/go-playground/validator/v10"
	cardHandler "github.com/jorgeloch/expenses-tracker/internal/handlers/card"
	debtorHandler "github.com/jorgeloch/expenses-tracker/internal/handlers/debtor"
	"github.com/jorgeloch/expenses-tracker/internal/handlers/owner"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
	"github.com/jorgeloch/expenses-tracker/internal/utils"
)

type Handler struct {
	OwnerHandler  *ownerHandler.Handler
	CardHandler   *cardHandler.Handler
	DebtorHandler *debtorHandler.Handler
}

func Init(s *service.Service) *Handler {
	validator := validator.New()
	err := utils.RegisterDateValidation(validator)
	if err != nil {
		log.Panic(err)
	}

	return &Handler{
		OwnerHandler:  ownerHandler.Init(s, validator),
		CardHandler:   cardHandler.Init(s, validator),
		DebtorHandler: debtorHandler.Init(s, validator),
	}
}
