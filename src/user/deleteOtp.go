package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func DeleteOtp(db *sqlx.DB, userId *uuid.UUID) error {
	if userId == nil {
		return fmt.Errorf("otpId not found")
	}

	_, err := db.Exec("DELETE FROM otps WHERE userId = $1", userId)

	if err != nil {
		return err
	}
	return nil
}
