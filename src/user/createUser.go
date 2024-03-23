package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

type User struct {
	Username string `db:"username"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func CreateUser(db *sqlx.DB, payload *models.RegisterPayload) error {

	userData := User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: encoder.HashPassword(payload.Password),
	}

	query := `INSERT INTO users (username,email,password) VALUES (:username,:email,:password) RETURNING *`

	_, err := db.NamedExec(query, userData)

	if err != nil {
		return err
	}

	return nil
}
