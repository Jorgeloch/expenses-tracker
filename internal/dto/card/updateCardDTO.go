package cardDTO

type UpdateCardDTO struct {
	Name         string `json:"name,omitempty"`
	Flag         string `json:"flag,omitempty"`
	DayOfClosing string `json:"day_of_closing,omitempty" validate:"omitempty,day"`
}
