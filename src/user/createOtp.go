package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateOtp(db *sqlx.DB, otp string, userId uuid.UUID) error {

	_, err := db.Exec("INSERT INTO otps(otp, userId) VALUES($1,$2) RETURNING id", otp, userId)
	if err != nil {
		return err
	}

	return nil
}
