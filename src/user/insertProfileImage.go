package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func InsertProfileImage(db *sqlx.DB, userId uuid.UUID, imageData []byte) (*uuid.UUID, error) {
	query := `INSERT INTO images(image,userId) VALUES($1, $2) RETURNING id`

	var imageId uuid.UUID
	err := db.QueryRowx(query, imageData, userId).Scan(&imageId)
	if err != nil {
		return nil, err
	}
	return &imageId, nil
}
