package user

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

// retives who is following the user.
func GetFollowerList(db *sqlx.DB, userId string) (*[]models.FollowerList, error) {
	var data []models.FollowerList

	err := db.Select(&data, "SELECT u.username AS username, u.id AS user_id, f.id AS id FROM followers f JOIN users u  ON f.follower_id = u.id WHERE followee_id = $1", userId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
