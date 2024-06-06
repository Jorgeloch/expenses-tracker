package debtorDTO

import "github.com/google/uuid"

type CreateDebtorDTO struct {
	Name    string    `json:"name" validate:"required"`
	OwnerID uuid.UUID `json:"owner_id" validate:"required,uuid"`
}
