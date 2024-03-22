package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserByEmail(db *sqlx.DB, email string) (*models.User, error) {
	users := models.User{}

	err := db.Get(&users, "Select * FROM users WHERE email = $1", email)

	if err != nil {

		return nil, err
	}

	return &users, nil

}
