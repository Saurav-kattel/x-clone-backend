package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func UpdateUsername(db *sqlx.DB, newUsername, userId string) error {
	if newUsername == "" {
		return fmt.Errorf("username not found")
	}

	_, err := db.Exec("UPDATE users SET username = $1 WHERE id = $2", newUsername, userId)
	if err != nil {
		return err
	}
	return nil
}
