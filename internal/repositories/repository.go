package repository

import (
	"github.com/jackc/pgx/v5"
	cardRepository "github.com/jorgeloch/expenses-tracker/internal/repositories/card"
	debtorRepository "github.com/jorgeloch/expenses-tracker/internal/repositories/debtor"
	"github.com/jorgeloch/expenses-tracker/internal/repositories/owner"
)

type Repository struct {
	OwnerRepository  *ownerRepository.Repository
	CardRepository   *cardRepository.Repository
	DebtorRepository *debtorRepository.Repository
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{
		OwnerRepository:  ownerRepository.Init(db),
		CardRepository:   cardRepository.Init(db),
		DebtorRepository: debtorRepository.Init(db),
	}
}
