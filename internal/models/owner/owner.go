package ownerModel

import "github.com/google/uuid"

type Owner struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid"`
	Email    string    `json:"email" validate:"required,email" errormgs:"Owner must have an email"`
	Password string    `json:"-" validate:"required,min=8" errormgs:"Owner must have a password and it must have at least 8 characters"`
	Name     string    `json:"name" validate:"required" error:"Owner must have a name"`
}
