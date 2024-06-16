package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetCoverImage(db *sqlx.DB, userId string) (*models.ProfileImage, error) {
	var data models.ProfileImage
	err := db.QueryRowx("SELECT id, image FROM cover_images WHERE user_id = $1", userId).StructScan(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
