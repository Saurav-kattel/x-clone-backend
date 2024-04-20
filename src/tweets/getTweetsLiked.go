package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweetLikedUser(db *sqlx.DB, tweetId string) (*[]models.LikedUsers, error) {
	var data []models.LikedUsers
	query := `SELECT users.username, likes.user_id as user_id FROM likes JOIN users ON likes.user_id = users.id WHERE likes.tweet_id = $1`

	err := db.Select(&data, query, tweetId)
	if err != nil {
		return nil, err
	}
	return &data, err
}
