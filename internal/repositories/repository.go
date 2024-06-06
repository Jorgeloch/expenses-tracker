package repository

import (
	"github.com/jackc/pgx/v5"
	cardRepository "github.com/jorgeloch/expenses-tracker/internal/repositories/card"
	"github.com/jorgeloch/expenses-tracker/internal/repositories/owner"
)

type Repository struct {
	OwnerRepository *ownerRepository.Repository
	CardRepository  *cardRepository.Repository
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{
		OwnerRepository: ownerRepository.Init(db),
		CardRepository:  cardRepository.Init(db),
	}
}
