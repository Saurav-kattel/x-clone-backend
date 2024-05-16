package user

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserByEmail(db *sqlx.DB, email string) (*models.User, error) {
	user := &models.User{}

	err := db.QueryRowx("SELECT * FROM users WHERE email = $1", email).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Role,
		&user.ImageId,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
	)

	if err != nil {
		log.Print("DB ERROR: " + err.Error())
		return nil, err
	}

	return user, nil
}
