package expenseModel

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Value       float64   `json:"value"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	UserID      int       `json:"user_id"`
	CardID      int       `json:"card_id"`
}
