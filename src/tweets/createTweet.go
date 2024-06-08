package tweets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func CreateTweets(db *sqlx.DB, payload *models.TweetsPayload, userId string) (string, error) {

	var id string
	if payload == nil || payload.Content == "" || userId == "" {
		return "", fmt.Errorf("create tweets: necessary data not provided userId = %v content = %+v", userId, payload)
	}
	err := db.QueryRowx("INSERT INTO tweets(content,userId) VALUES($1,$2) RETURNING id", payload.Content, userId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
