package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetComments(db *sqlx.DB, postId string, pageNumber, pageSize int) (*[]models.CommentData, error) {
	var data []models.CommentData
	offset := (pageNumber - 1) * pageSize
	err := db.Select(&data, `SELECT comments.*, users.username AS username 
		FROM comments JOIN users ON comments.user_id = users.id
		WHERE tweet_id = $1
		ORDER BY created_at DESC
		LIMIT $2
		OFFSET $3
		`, postId, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
