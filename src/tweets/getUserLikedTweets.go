package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserLikedTweets(db *sqlx.DB, pageNumber, pageSize int, userId string) (*[]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT DISTINCT t.id, t.content, t.imageid, t.created_at, t.updated_at,
        u.id as userid, u.username as author_username,
        CONCAT(u.first_name,' ',u.last_name) as author
        FROM tweets t
        LEFT JOIN users u ON t.userid = u.id
	LEFT JOIN likes l ON l.tweet_id = t.id
	WHERE l.user_id = $1
        ORDER BY t.created_at DESC
        LIMIT $2 OFFSET $3 
	`

	err := db.Select(&tweets, query, userId, pageSize, offset)
	if err != nil {
		return nil, err
	}

	return &tweets, nil

}
