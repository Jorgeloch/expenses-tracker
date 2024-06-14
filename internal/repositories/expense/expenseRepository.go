package expenseRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	expenseModel "github.com/jorgeloch/expenses-tracker/internal/models/expense"
)

type Repository struct {
	DB *pgx.Conn
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll(ownerID string) ([]expenseModel.Expense, error) {
	rows, err := r.DB.Query(context.Background(), `
    SELECT * FROM expenses
    WHERE owner_id=$1
    `, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var expenses []expenseModel.Expense
	for rows.Next() {
		var expense expenseModel.Expense
		err := rows.Scan(
			&expense.ID,
			&expense.Value,
			&expense.Description,
			&expense.Date,
			&expense.OwnerID,
			&expense.DebtorID,
			&expense.CardID,
		)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}
	return expenses, nil
}

func (r *Repository) GetByID(ownerID string, id string) (expenseModel.Expense, error) {
	var expense expenseModel.Expense

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM expenses
    WHERE id=$1
    AND owner_id=$2
    `, id, ownerID).Scan(
		&expense.ID,
		&expense.Value,
		&expense.Description,
		&expense.Date,
		&expense.OwnerID,
		&expense.DebtorID,
		&expense.CardID,
	)

	if err != nil {
		return expenseModel.Expense{}, err
	}

	return expense, err
}

func (r *Repository) Create(expense expenseModel.Expense) error {
	args := pgx.NamedArgs{
		"id":          expense.ID,
		"value":       expense.Value,
		"description": expense.Description,
		"date":        expense.Date,
		"owner_id":    expense.OwnerID,
		"debtor_id":   expense.DebtorID,
		"card_id":     expense.CardID,
	}

	_, err := r.DB.Exec(context.Background(),
		`
    INSERT INTO expenses (id, value, description, date, owner_id, debtor_id, card_id)
    VALUES (@id, @value, @description, @date, @owner_id, @debtor_id, @card_id)
    `,
		args)

	return err
}

func (r *Repository) Update(expense expenseModel.Expense) error {
	args := pgx.NamedArgs{
		"id":          expense.ID,
		"value":       expense.Value,
		"description": expense.Description,
		"date":        expense.Date,
		"owner_id":    expense.OwnerID,
		"debtor_id":   expense.DebtorID,
		"card_id":     expense.CardID,
	}
	_, err := r.DB.Exec(context.Background(),
		`
    UPDATE expenses
    SET value=@value, description=@description, date=@date, card_id=@card_id, debtor_id=@debtor_id
    WHERE id=@id
    AND owner_id=@owner_id
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
    DELETE FROM expenses 
    WHERE id=$1
    AND owner_id=$2
    `,
		id, ownerID)
	return err
}
