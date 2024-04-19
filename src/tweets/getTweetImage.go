package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweetImage(db *sqlx.DB, imgId string) (*models.TweetImage, error) {
	var data models.TweetImage

	err := db.QueryRowx("SELECT image FROM tweetsimages WHERE id = $1", imgId).Scan(&data.Image)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
