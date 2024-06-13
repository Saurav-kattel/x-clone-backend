package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserComment(db *sqlx.DB, userId, tweetId string, pageSize, pageNumber int) (*[]models.CommentData, error) {

	var data []models.CommentData

	offset := (pageNumber - 1) * pageSize
	err := db.Select(&data, `
	SELECT DISTINCT comments.*,users.username as username FROM comments
	JOIN users ON comments.user_id = users.id
	LEFT JOIN reply r ON r.comment_id = comments.id
	WHERE (comments.user_id = $1 AND comments.tweet_id = $2) OR (r.replied_from = $1 OR r.replied_to = $1)
	ORDER by created_at DESC LIMIT $3 OFFSET $4`, userId, tweetId, pageSize, offset)

	if err != nil {
		return nil, err
	}

	return &data, nil
}
