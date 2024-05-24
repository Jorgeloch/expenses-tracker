package repository

import (
	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/internal/repositories/owner"
)

type Repository struct {
	OwnerRepository *ownerRepository.Repository
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{
		OwnerRepository: ownerRepository.Init(db),
	}
}
