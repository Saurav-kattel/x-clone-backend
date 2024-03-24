package user

import "github.com/jmoiron/sqlx"

func DeleteAccount(db *sqlx.DB, userId string) error {

	query := `DELETE FROM users WHERE id = $1`

	_, err := db.Exec(query, userId)
	if err != nil {
		return err
	}

	return nil
}
