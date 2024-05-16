package user

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserByUsername(db *sqlx.DB, username string) (*models.User, error) {
	user := &models.User{}

	err := db.QueryRowx("SELECT * FROM users WHERE username = $1", username).Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.Username,
		&user.Role,
		&user.ImageId,
		&user.CreatedAt,
		&user.LastName,
		&user.FirstName,
	)

	if err != nil {
		log.Print("DB ERROR: " + err.Error())
		return nil, err
	}

	return user, nil
}
