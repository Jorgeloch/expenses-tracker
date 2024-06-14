package expenseDTO

import (
	"time"

	"github.com/google/uuid"
)

type CreateExpenseDTO struct {
	Value       float64   `json:"value" validate:"reqired,gt=0" errormgs:"Value must be grather than 0"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" validate:"required,datetime" errormgs:"Expense must have a date"`
	DebtorID    uuid.UUID `json:"debtor_id" validate:"uuid"`
	CardID      uuid.UUID `json:"card_id" validate:"required"`
}
