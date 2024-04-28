package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetFollowerList(db *sqlx.DB, userId string) (*[]models.FollowerList, error) {
	var data []models.FollowerList

	err := db.Select(&data, "SELECT u.username as username, u.id as user_id, f.id as id FROM followers f join users u  ON f.followee_id = u.id where follower_id = $1", userId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
