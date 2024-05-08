package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetComments(db *sqlx.DB, postId string) (*[]models.CommentData, error) {
	var data []models.CommentData

	err := db.Select(&data, "SELECT * FROM comments WHERE tweet_id = $1", postId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
