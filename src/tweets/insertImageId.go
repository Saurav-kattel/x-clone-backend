package tweets

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InsertImageId(db *sqlx.DB, tweetId, imageId string) error {
	if tweetId == "" || imageId == "" {
		return fmt.Errorf("insertImageId: necessary data not found")
	}
	_, err := db.Exec("UPDATE tweets SET imageId =  $1 WHERE id = $2", imageId, tweetId)
	if err != nil {
		return err
	}
	return nil
}
