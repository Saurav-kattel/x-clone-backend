package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func InsertImageId(db *sqlx.DB, imageId string, userId string) error {

	if imageId == "" || userId == "" {
		return fmt.Errorf("invalid params")
	}
	//query to insert imageId where user.id equals email
	_, err := db.Exec("UPDATE users SET image_id = $1 WHERE id = $2", imageId, userId)

	if err != nil {
		return err
	}
	return nil
}
