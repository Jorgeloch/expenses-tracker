package expenseModel

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          int       `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	OwnerID     uuid.UUID `json:"owner_id"`
	CardID      uuid.UUID `json:"card_id"`
}
