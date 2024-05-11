package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetFirstReply(db *sqlx.DB, tweetId, commentId string) (*[]models.ReplyData, error) {
	var data []models.ReplyData

	query := `
		SELECT
    r.id, 
    r.reply,
    r.tweet_id, 
    r.replied_to, r.replied_from, r.parent_id,
    r.created_at,
    r.comment_id,
		tu.username AS replied_to_username,
		fu.username AS replied_from_username
		FROM reply r
		JOIN users tu ON r.replied_to = tu.id
		JOIN users fu ON r.replied_from = fu.id
		WHERE r.tweet_id = $1
    AND r.comment_id = $2
    AND r.parent_id IS NULL
		ORDER BY r.id`

	err := db.Select(&data, query, tweetId, commentId)
	if err != nil {

		return nil, err
	}
	return &data, nil
}
