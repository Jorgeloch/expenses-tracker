package expenseDTO

import (
	"time"

	"github.com/google/uuid"
)

type UpdateExpenseDTO struct {
	Value       float64   `json:"value,omitempty" validate:"omitempty,gt=0"`
	Description string    `json:"description,omitempty"`
	Date        time.Time `json:"date,omitempty" validate:"omitempty,datetime"`
	DebtorID    uuid.UUID `json:"debtor_id,omitempty" validate:"omitempty,uuid"`
	CardID      uuid.UUID `json:"card_id,omitempty"`
}
