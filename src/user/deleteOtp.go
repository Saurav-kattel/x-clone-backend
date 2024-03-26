package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DeleteOtp(db *sqlx.DB, userId string) error {
	if userId == "" {
		return fmt.Errorf("otpId not found")
	}

	_, err := db.Exec("DELETE FROM otps WHERE userId = $1", userId)

	if err != nil {
		return err
	}
	return nil
}
