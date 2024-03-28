package tweets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DeleteTweet(db *sqlx.DB, tweetId, userId string) error {
	if tweetId == "" {
		return fmt.Errorf("tweetId not found")
	}

	_, err := db.Exec("DELETE FROM tweets WHERE id = $1 AND userId = $2", tweetId, userId)
	if err != nil {
		return err
	}
	return nil
}
