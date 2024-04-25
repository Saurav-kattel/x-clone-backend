package user

import "github.com/jmoiron/sqlx"

func HandleUnFollow(db *sqlx.DB, followerId, followeeId string) error {
	_, err := db.Exec("DELETE FROM followers WHERE follower_id = $1 AND followee_id = $2", followerId, followeeId)
	if err != nil {
		return err
	}
	return nil
}
