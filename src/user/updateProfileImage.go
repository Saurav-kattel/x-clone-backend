package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

func UpdateProfileImage(db *sqlx.DB, userId uuid.UUID, imageData []byte) (*uuid.UUID, error) {
	query := `UPDATE  images SET image = $1 WHERE userId = $2 RETURNING id`

	var imageId uuid.UUID
	err := db.QueryRowx(query, imageData, userId).Scan(&imageId)
	if err != nil {
		return nil, err
	}
	return &imageId, nil
}
