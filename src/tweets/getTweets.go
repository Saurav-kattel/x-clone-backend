package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweets(db *sqlx.DB, pageNumber, pageSize int) ([]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT t.id, t.content, t.imageid, t.created_at, t.updated_at,
        u.id as userid, u.username as author
        FROM tweets t
        JOIN users u ON t.userid = u.id
        ORDER BY t.created_at DESC
        LIMIT $1 OFFSET $2 
    `

	err := db.Select(&tweets, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return tweets, nil
}
