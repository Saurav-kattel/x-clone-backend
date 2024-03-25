package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func UpdateForgottenPassword(db *sqlx.DB, hash, userId string) error {

	if hash == "" {
		return fmt.Errorf("password not found")
	}

	if userId == "" {
		return fmt.Errorf("userId not found")
	}

	_, err := db.Exec("UPDATE users SET password = $1 WHERE email = $2", hash, userId)

	if err != nil {
		return err
	}
	return nil
}
