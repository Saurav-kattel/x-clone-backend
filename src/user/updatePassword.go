package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func UpdatePassword(db *sqlx.DB, userId, hash string) error {

	if userId == "" {
		return fmt.Errorf("user id not found")
	}

	_, err := db.Exec("UPDATE users SET password = $1 WHERE id = $2", hash, userId)
	if err != nil {
		return err
	}

	return nil
}
