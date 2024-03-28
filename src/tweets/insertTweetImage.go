package tweets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InsertTweetImage(db *sqlx.DB, imageData []byte, tweetId string) (string, error) {

	var id string
	if len(imageData) == 0 || tweetId == "" {
		return "", fmt.Errorf("insert image: necessary data not found")
	}

	err := db.QueryRowx("INSERT INTO tweetsimages(image, tweetId) VALUES($1,$2) RETURNING id", imageData, tweetId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
