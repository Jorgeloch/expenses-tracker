package cardDTO

type CreateCardDTO struct {
	Name         string `json:"name" validate:"required"`
	Flag         string `json:"flag" validate:"required"`
	DayOfClosing string `json:"day_of_closing" validate:"required,day"`
}
