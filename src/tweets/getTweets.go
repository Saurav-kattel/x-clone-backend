package tweets

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweets(db *sqlx.DB, pageNumber, pageSize int, vis string) ([]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT t.id, t.content, t.imageid, t.created_at, t.updated_at,
        u.id as userid, u.username as author_username,
        CONCAT(u.first_name,' ',u.last_name) as author
        FROM tweets t
        JOIN users u ON t.userid = u.id
	WHERE t.visibility = $3
        ORDER BY t.created_at DESC
        LIMIT $1 OFFSET $2 `

	err := db.Select(&tweets, query, pageSize, offset, vis)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return tweets, nil
}
