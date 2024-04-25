package user

import "github.com/jmoiron/sqlx"

func HandleFollow(db *sqlx.DB, followerId, followeeId string) error {

	_, err := db.Exec("INSERT INTO followers(follower_id, followee_id) VALUES($1,$2)", followerId, followeeId)
	if err != nil {
		return err
	}
	return nil
}
