package tweets

import "github.com/jmoiron/sqlx"

func HandleUnlike(db *sqlx.DB, tweetId, userId string) error {
	_, err := db.Exec("DELETE FROM likes WHERE tweet_id = $1 AND user_id = $2", tweetId, userId)
	if err != nil {
		return err
	}
	return nil
}
