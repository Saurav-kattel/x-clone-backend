package tweets

import "github.com/jmoiron/sqlx"

func CreateComment(db *sqlx.DB, comment, userId, tweetId string) error {
	_, err := db.Exec("INSERT INTO comments(comment,user_id,tweet_id) VALUES($1,$2,$3)", comment, userId, tweetId)
	if err != nil {
		return err
	}
	return nil
}
