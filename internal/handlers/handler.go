package handlers

import (
	"github.com/jorgeloch/expenses-tracker/internal/handlers/owner"
	service "github.com/jorgeloch/expenses-tracker/internal/services"
)

type Handler struct {
	OwnerHandler *ownerHandler.Handler
}

func Init(s *service.Service) *Handler {
	return &Handler{
		OwnerHandler: ownerHandler.Init(s),
	}
}
