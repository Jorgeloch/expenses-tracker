package debtorRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	debtorModel "github.com/jorgeloch/expenses-tracker/internal/models/debtor"
)

type Repository struct {
	DB *pgx.Conn
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll(ownerID string) ([]debtorModel.Debtor, error) {
	rows, err := r.DB.Query(context.Background(), `
    SELECT * FROM debtors
    WHERE owner_id=$1
    `, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var debtors []debtorModel.Debtor
	for rows.Next() {
		var debtor debtorModel.Debtor
		err := rows.Scan(&debtor.ID, &debtor.Name, &debtor.OwnerID)
		if err != nil {
			return nil, err
		}
		debtors = append(debtors, debtor)
	}
	return debtors, nil
}

func (r *Repository) GetByID(ownerID string, id string) (debtorModel.Debtor, error) {
	var debtor debtorModel.Debtor

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM debtors
    WHERE id=$1
    AND owner_id=$2
    `, id, ownerID).Scan(&debtor.ID, &debtor.Name, &debtor.OwnerID)

	if err != nil {
		return debtorModel.Debtor{}, err
	}

	return debtor, err
}

func (r *Repository) Create(debtor debtorModel.Debtor) error {
	args := pgx.NamedArgs{
		"id":       debtor.ID,
		"name":     debtor.Name,
		"owner_id": debtor.OwnerID,
	}

	_, err := r.DB.Exec(context.Background(),
		`
    INSERT INTO debtors (id, name, owner_id)
    VALUES (@id, @name, @owner_id)
    `,
		args)

	return err
}

func (r *Repository) Update(debtor debtorModel.Debtor) error {
	args := pgx.NamedArgs{
		"id":       debtor.ID,
		"name":     debtor.Name,
		"owner_id": debtor.OwnerID,
	}
	_, err := r.DB.Exec(context.Background(),
		`
    UPDATE debtors
    SET name=@name, owner_id=@owner_id
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
    DELETE FROM debtors 
    WHERE id=$1
    AND owner_id=$2
    `,
		id, ownerID)
	return err
}
