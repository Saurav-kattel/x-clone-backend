package user

import "github.com/jmoiron/sqlx"

func CreateCoverImage(db *sqlx.DB, userId string, image []byte) error {
	_, err := db.Exec("INSERT INTO cover_images(user_id,image) VALUES($1,$2)", userId, image)
	return err
}
