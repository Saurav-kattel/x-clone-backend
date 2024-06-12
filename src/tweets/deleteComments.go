package tweets

import "github.com/jmoiron/sqlx"

func DeleteComment(db *sqlx.DB, commentId string) error {
	_, err := db.Exec("DELETE FROM comments WHERE id = $1", commentId)
	return err
}

func DeleteReply(db *sqlx.DB, replyId string) error {
	_, err := db.Exec("DELETE FROM reply WHERE id = $1", replyId)
	return err
}
