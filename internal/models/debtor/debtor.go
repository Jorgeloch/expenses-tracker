package debtorModel

import "github.com/google/uuid"

type Debtor struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid"`
	Name    string    `json:"name" validate:"required"`
	OwnerID uuid.UUID `json:"owner_id" validate:"required,uuid"`
}
