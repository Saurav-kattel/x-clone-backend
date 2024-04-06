package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetProfileImage(db *sqlx.DB, userId string, imageId any) (*models.ProfileImage, error) {
	data := &models.ProfileImage{}

	err := db.QueryRowx("SELECT id,image FROM images WHERE userId = $1 AND id = $2", userId, imageId).StructScan(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
