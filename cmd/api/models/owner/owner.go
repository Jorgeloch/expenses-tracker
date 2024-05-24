package ownerModel

import "github.com/google/uuid"

type Owner struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"-" validate:"required,min=8"`
	Name     string    `json:"name" validate:"required"`
}
