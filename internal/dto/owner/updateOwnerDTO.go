package ownerDTO

type UpdateOwnerDTO struct {
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=8" errormgs:"Password must have at least 8 characters"`
	Name     string `json:"name,omitempty"`
}
