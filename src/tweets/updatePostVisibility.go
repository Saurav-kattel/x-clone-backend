package tweets

import "github.com/jmoiron/sqlx"

func UpdatePostVisibility(db *sqlx.DB, tweet_id, vis, userId string) error {
	_, err := db.Exec("UPDATE tweets SET visibility = $1 WHERE id = $2 AND userId = $3", vis, tweet_id, userId)
	return err
}
