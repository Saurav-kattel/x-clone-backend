package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func SearchResult(db *sqlx.DB, query string) (*[]models.UserData, error) {

	var data []models.UserData

	err := db.Select(&data, `
		SELECT 
		id, username, first_name, last_name, role, created_at, email, image_id
		FROM users WHERE username ILIKE $1
	`, query)

	if err != nil {
		return nil, err
	}
	return &data, err

}
