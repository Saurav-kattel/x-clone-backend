package tweets

import "github.com/jmoiron/sqlx"

func CreateReplies(db *sqlx.DB, reply, repliedFrom, tweetId string, parent_comment_id *string, repliedTo string) error {
	_, err := db.Exec("INSERT INTO reply(reply,replied_to,replied_from,tweet_id, parent_id) VALUES($1,$2,$3,$4,$5)", reply, repliedTo, repliedFrom, tweetId, parent_comment_id)
	if err != nil {
		return err
	}
	return nil
}
