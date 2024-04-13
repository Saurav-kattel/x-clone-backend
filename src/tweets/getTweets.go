package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweets(db *sqlx.DB, pageNumber, pageSize int) (*[]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize

	err := db.Select(&tweets, "SELECT * FROM tweets LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		return nil, err
	}
	return &tweets, nil
}
