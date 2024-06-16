package user

import "github.com/jmoiron/sqlx"

func UpdateCoverImage(db *sqlx.DB, userId string, image []byte) error {

	query := `UPDATE  cover_images SET image = $1 WHERE user_Id = $2 RETURNING id`

	_, err := db.Exec(query, image, userId)
	return err
}
