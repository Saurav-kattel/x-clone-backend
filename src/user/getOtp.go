package user

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Data struct {
	Id        string    `db:"id"`
	Otp       string    `db:"otp"`
	UserId    string    `db:"userId"`
	ExpiresAt time.Time `db:"expires_at"`
	Email     string    `db:"email"`
}

func GetOtpWithUser(db *sqlx.DB, userId string) (*Data, error) {
	var data Data

	query := `SELECT 
    o.id AS otp_id, 
    o.otp, 
    o.userId, 
    o.expires_at, 
    u.email AS email
	FROM
    otps o
	JOIN
    users u ON o.userId = u.id
	WHERE
    o.userId = $1`

	err := db.QueryRowx(query, userId).Scan(
		&data.Id,
		&data.Otp,
		&data.UserId,
		&data.ExpiresAt,
		&data.Email,
	)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
