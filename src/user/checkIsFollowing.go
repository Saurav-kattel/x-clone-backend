package user

import (
	"github.com/jmoiron/sqlx"
)

func IsFollowing(db *sqlx.DB, followerId, followeeId string) error {
	var id string
	err := db.QueryRowx("SELECT id FROM followers WHERE follower_id = $1 AND followee_id = $2", followerId, followeeId).Scan(&id)

	if err != nil {
		return err
	}
	return nil
}
