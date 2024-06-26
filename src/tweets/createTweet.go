package tweets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func CreateTweets(db *sqlx.DB, payload *models.TweetsPayload, userId string) (string, error) {

	var id string
	if payload == nil || userId == "" {
		return "", fmt.Errorf("create tweets: necessary data not provided userId = %v content = %+v", userId, payload)
	}
	err := db.QueryRowx("INSERT INTO tweets(content,userId, visibility) VALUES($1,$2,$3) RETURNING id", payload.Content, userId, payload.Visibility).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
