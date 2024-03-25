package user

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetOtp(db *sqlx.DB, userId *uuid.UUID) (*models.Otp, error) {
	var data models.Otp

	if userId == nil {
		return nil, fmt.Errorf("user id not found")
	}

	err := db.QueryRowx("SELECT * FROM otps WHERE userId = $1", userId).Scan(&data.ID, &data.Otp, &data.UserId, &data.CreatedAt, &data.ExpiresAt)

	if err != nil {
		return nil, err
	}
	return &data, nil
}
