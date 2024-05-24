package ownerRepository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jorgeloch/expenses-tracker/cmd/api/models/owner"
)

type Repository struct {
	DB *pgx.Conn
}

func Init(db *pgx.Conn) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]ownerModel.Owner, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []ownerModel.Owner
	for rows.Next() {
		var user ownerModel.Owner
		err := rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *Repository) GetByID(id int) (ownerModel.Owner, error) {
	var user ownerModel.Owner

	err := r.DB.QueryRow(context.Background(),
		`
    SELECT * FROM users
    WHERE id=$1
    `, id).Scan(&user.ID, &user.Name)

	if err != nil {
		return user, err
	}

	return user, err
}

func (r *Repository) Create(user ownerModel.Owner) (uuid.UUID, error) {
	args := pgx.NamedArgs{
		"id":   user.ID,
		"name": user.Name,
	}

	_, err := r.DB.Exec(context.Background(),
		`
    INSERT INTO users (id, name)
    VALUES (@id, @name)
    `,
		args)

	if err != nil {
		return uuid.Nil, err
	}

	return user.ID, nil
}

func (r *Repository) Update(user ownerModel.Owner) error {
	args := pgx.NamedArgs{
		"id":   user.ID,
		"name": user.Name,
	}
	_, err := r.DB.Exec(context.Background(),
		`
    UPDATE users
    SET name=@userName
    WHERE id=@userID
    `,
		args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	_, err := r.DB.Exec(context.Background(),
		`
    DELETE FROM users WHERE id=$1
    `,
		id)
	return err
}
