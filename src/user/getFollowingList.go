package user

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

// retrives  following
func GetFollowingList(db *sqlx.DB, userName string) (*[]models.FollowerList, error) {
	var data []models.FollowerList
	err := db.Select(&data, "SELECT u.username AS username, u.id AS user_id, f.id AS id FROM followers f JOIN users u  ON f.followee_id = u.id WHERE follower_id = ( SELECT id FROM users WHERE username = $1 )", userName)
	if err != nil {
		return nil, err
	}

	log.Print("twere", data)
	return &data, nil
}
