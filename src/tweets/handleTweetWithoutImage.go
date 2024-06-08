package tweets

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func HandleTweetsWithoutImage(w http.ResponseWriter, db *sqlx.DB, data *models.TweetsPayload, userId string) error {
	_, err := CreateTweets(db, data, userId)
	return err
}
