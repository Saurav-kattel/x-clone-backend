package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweetLikedUser(db *sqlx.DB, tweetId string) (*[]models.LikedUsers, error) {
	var data []models.LikedUsers
	err := db.Select(&data, "SELECT u.username, l.id, l.tweet_id FROM user u JOIN likes l ON u.id = l.user_id WHERE l.tweet_id = $1", tweetId)
	if err != nil {
		return nil, err
	}
	return &data, err
}
