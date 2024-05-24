package debtorModel

import "github.com/google/uuid"

type Debtor struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	OwnerID uuid.UUID `json:"owner_id"`
}
