package cardModel

import "github.com/google/uuid"

type Card struct {
	ID      uuid.UUID `json:"id" validate:"required,uuid"`
	Name    string    `json:"name" validate:"required"`
	Flag    string    `json:"flag" validate:"required"`
	OwnerID uuid.UUID `json:"owner_id" validate:"required,uuid"`
}
