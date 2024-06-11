package ownerRepository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/internal/models/owner"
)

type Repository struct {
	DB *pgx.Conn
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]ownerModel.Owner, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT * FROM owners")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var owners []ownerModel.Owner
	for rows.Next() {
		var owner ownerModel.Owner
		err := rows.Scan(&owner.ID, &owner.Name, &owner.Email, &owner.Password)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}
	return owners, nil
}

func (r *Repository) GetByID(id string) (ownerModel.Owner, error) {
	var owner ownerModel.Owner

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM owners
    WHERE id=$1
    `, id).Scan(&owner.ID, &owner.Name, &owner.Email, &owner.Password)

	if err != nil {
		return owner, err
	}

	return owner, err
}

func (r *Repository) GetByEmail(email string) (ownerModel.Owner, error) {
	var owner ownerModel.Owner

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM owners
    WHERE email=$1
    `, email).Scan(&owner.ID, &owner.Name, &owner.Email, &owner.Password)

	if err != nil {
		return owner, err
	}

	return owner, err
}

func (r *Repository) Create(owner ownerModel.Owner) error {
	args := pgx.NamedArgs{
		"id":       owner.ID,
		"name":     owner.Name,
		"email":    owner.Email,
		"password": owner.Password,
	}

	_, err := r.DB.Exec(context.Background(),
		`
    INSERT INTO owners (id, name, email, password)
    VALUES (@id, @name, @email, @password)
    `,
		args)

	return err
}

func (r *Repository) Update(owner ownerModel.Owner) error {
	args := pgx.NamedArgs{
		"id":       owner.ID,
		"name":     owner.Name,
		"email":    owner.Email,
		"password": owner.Password,
	}
	_, err := r.DB.Exec(context.Background(),
		`
    UPDATE owners
    SET name=@name, email=@email, password=@password
    WHERE id=@id
    `,
		args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.DB.Exec(context.Background(),
		`
    DELETE FROM owners WHERE id=$1
    `,
		id)
	return err
}
