package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

type User struct {
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func CreateUser(db *sqlx.DB, payload *models.RegisterPayload) (*string, error) {
	var id string

	userData := User{
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  encoder.HashPassword(payload.Password),
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
	}

	query := `INSERT INTO users (username,email,password,first_name,last_name) VALUES (:username,:email,:password,:first_name,:last_name) RETURNING id`

	rows, err := db.NamedQuery(query, userData)
	scanErr := rows.Scan(&id)

	if scanErr != nil {
		return nil, scanErr
	}

	if err != nil {
		return nil, err
	}

	return &id, nil
}
