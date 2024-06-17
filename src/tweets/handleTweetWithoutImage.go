package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func HandleTweetsWithoutImage(db *sqlx.DB, data *models.TweetsPayload, userId string) error {
	_, err := CreateTweets(db, data, userId)
	return err
}
