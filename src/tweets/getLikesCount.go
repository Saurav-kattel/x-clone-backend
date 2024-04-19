package tweets

import "github.com/jmoiron/sqlx"

func GetLikesCount(db *sqlx.DB, tweetId string) (int, error) {
	var likes int

	err := db.QueryRowx("SELECT COUNT(*) FROM likes WHERE tweet_id = $1", tweetId).Scan(&likes)
	if err != nil {
		return -1, err
	}
	return likes, nil
}
