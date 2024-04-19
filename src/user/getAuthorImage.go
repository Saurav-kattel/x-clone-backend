package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetAuthorImage(db *sqlx.DB, userId string) (*models.AuthorImage, error) {
	data := models.AuthorImage{}

	err := db.QueryRowx("SELECT image FROM images WHERE userId = $1", userId).Scan(&data.Image)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
