package ownerDTO

type CreateOwnerDTO struct {
	Email    string `json:"email" validate:"required,email" errormgs:"Owner must have an email"`
	Password string `json:"password" validate:"required,min=8"`
	Name     string `json:"name" validate:"required"`
}
