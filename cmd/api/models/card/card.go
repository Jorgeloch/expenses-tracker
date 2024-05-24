package cardModel

import "github.com/google/uuid"

type Card struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Flag    string    `json:"flag"`
	OwnerID uuid.UUID `json:"owner_id"`
}
