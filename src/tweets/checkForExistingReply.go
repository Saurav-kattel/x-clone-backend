package tweets

import "github.com/jmoiron/sqlx"

func CheckForExistingReply(db *sqlx.DB, replyId string) error {
	var id string
	err := db.QueryRowx("SELECT id FROM reply WHERE id = $1", replyId).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
