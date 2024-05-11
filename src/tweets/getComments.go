package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetComments(db *sqlx.DB, postId string) (*[]models.CommentData, error) {
	var data []models.CommentData

	// Retrieve comments for the given postId, ordering them by creation time in descending order
	err := db.Select(&data, "SELECT comments.*,users.username as username FROM comments JOIN users ON comments.user_id = user_id WHERE tweet_id = $1 ORDER BY created_at DESC", postId)
	if err != nil {
		return nil, err
	}
	// Return the pointer to the slice of comments and nil for the error
	return &data, nil
}
