package tweets

import "github.com/jmoiron/sqlx"

func CreateReplies(db *sqlx.DB, comment, userId, tweetId, parent_comment_id string) error {
	_, err := db.Exec("INSERT INTO comments(comment,user_id,tweet_id, parent_comment_id) VALUES($1,$2,$3,&4)", comment, userId, tweetId, parent_comment_id)
	if err != nil {
		return err
	}
	return nil
}
