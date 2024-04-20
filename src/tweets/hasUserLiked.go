package tweets

import (
	"github.com/jmoiron/sqlx"
)

func HasUserLiked(db *sqlx.DB, userId, tweetId string) error {
	var id string
	err := db.QueryRowx("SELECT id FROM likes WHERE tweet_id = $1 AND user_id = $2", tweetId, userId).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
