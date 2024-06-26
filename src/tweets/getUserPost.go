package tweets

import (
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetUserPost(db *sqlx.DB, pageSize, pageNumber int, username, vis string) (*[]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize
	query := `
  SELECT 
    t.id, 
    t.content, 
    t.imageid, 
    t.created_at, 
    t.updated_at,
    t.visibility,
    u.id AS userid, 
    u.username AS author_username,

    CONCAT(u.first_name, ' ', u.last_name) AS author
FROM 
    tweets t
JOIN 
    users u 
ON 
    t.userid = u.id
WHERE 
    u.username = $1 
AND
  t.visibility = $4
ORDER BY 
    t.created_at DESC
LIMIT 
    $2
OFFSET $3;
    `

	err := db.Select(&tweets, query, username, pageSize, offset, vis)
	if err != nil {
		return nil, err
	}

	return &tweets, nil
}
