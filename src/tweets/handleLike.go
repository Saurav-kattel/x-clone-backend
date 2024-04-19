package tweets

import "github.com/jmoiron/sqlx"

func HandleLike(db *sqlx.DB, tweetId, userId string) error {
	_, err := db.Exec("INSERT INTO likes(user_id,tweet_id) VALUES($1,$2)", userId, tweetId)
	if err != nil {
		return err
	}
	return nil
}
