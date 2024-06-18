package debtorDTO

type CreateDebtorDTO struct {
	Name string `json:"name" validate:"required"`
}
