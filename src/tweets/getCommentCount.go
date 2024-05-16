package tweets

import "github.com/jmoiron/sqlx"

func GetCommentCount(db *sqlx.DB, tweetId string) (*int, error) {
	var count int
	err := db.QueryRowx("SELECT COUNT(id) FROM comments WHERE tweet_id = $1", tweetId).Scan(&count)
	if err != nil {
		return nil, err
	}
	return &count, nil

}
