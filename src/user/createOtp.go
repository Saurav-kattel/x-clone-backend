package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func CreateOtp(db *sqlx.DB, otp string, userId uuid.UUID) (*uuid.UUID, error) {

	var otpId *uuid.UUID

	err := db.QueryRowx("INSERT INTO otps(otp, userId) VALUES($1,$2) RETURNING id", otp, userId).Scan(&otpId)
	if err != nil {
		return nil, err
	}

	return otpId, nil
}
