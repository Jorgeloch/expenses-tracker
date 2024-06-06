package cardRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	cardModel "github.com/jorgeloch/expenses-tracker/internal/models/card"
)

type Repository struct {
	DB *pgx.Conn
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll(ownerID string) ([]cardModel.Card, error) {
	rows, err := r.DB.Query(context.Background(), `
    SELECT * FROM cards
    WHERE owner_id=$1
    `, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cards []cardModel.Card
	for rows.Next() {
		var card cardModel.Card
		err := rows.Scan(&card.ID, &card.Name, &card.Flag, &card.OwnerID)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func (r *Repository) GetByID(ownerID string, id string) (cardModel.Card, error) {
	var card cardModel.Card

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM cards
    WHERE id=$1
    AND owner_id=$2
    `, id, ownerID).Scan(&card.ID, &card.Name, &card.Flag, &card.OwnerID)

	if err != nil {
		return cardModel.Card{}, err
	}

	return card, err
}

func (r *Repository) Create(card cardModel.Card) error {
	args := pgx.NamedArgs{
		"id":       card.ID,
		"name":     card.Name,
		"flag":     card.Flag,
		"owner_id": card.OwnerID,
	}

	_, err := r.DB.Exec(context.Background(),
		`
    INSERT INTO cards (id, name, flag, owner_id)
    VALUES (@id, @name, @flag, @owner_id)
    `,
		args)

	return err
}

func (r *Repository) Update(card cardModel.Card) error {
	args := pgx.NamedArgs{
		"id":       card.ID,
		"name":     card.Name,
		"flag":     card.Flag,
		"owner_id": card.OwnerID,
	}
	_, err := r.DB.Exec(context.Background(),
		`
    UPDATE cards
    SET name=@name, flag=@flag, owner_id=@owner_id
    WHERE id=@id
    `,
		args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(ownerID string, id string) error {
	_, err := r.DB.Exec(context.Background(),
		`
    DELETE FROM cards 
    WHERE id=$1
    AND owner_id=$2
    `,
		id, ownerID)
	return err
}
