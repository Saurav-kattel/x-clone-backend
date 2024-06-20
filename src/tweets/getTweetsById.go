package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweetsById(db *sqlx.DB, tweetId, vis string) (*models.Tweets, error) {
	tweets := models.Tweets{}

	query := `
        SELECT t.id, t.visibility,t.content, t.imageid, t.created_at, t.updated_at,
        u.id as userid, u.username as author_username,
        CONCAT(u.first_name,' ',u.last_name) as author
        FROM tweets t
        JOIN users u ON t.userid = u.id
        WHERE t.id = $1 AND  t.visibility = $2
        ORDER BY t.created_at DESC
    `

	err := db.QueryRowx(query, tweetId, vis).StructScan(&tweets)
	if err != nil {
		return nil, err
	}
	return &tweets, nil
}
