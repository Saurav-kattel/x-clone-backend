package tweets

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func GetTweets(db *sqlx.DB, pageNumber, pageSize int, vis, userId string) ([]models.Tweets, error) {
	tweets := []models.Tweets{}
	offset := (pageNumber - 1) * pageSize

	query := `
        SELECT DISTINCT t.id, t.content, t.visibility, t.imageid, t.created_at, t.updated_at,  u.id as userid, u.username as author_username,
	CONCAT(u.first_name,' ',u.last_name) as author
           FROM tweets t
           JOIN users u ON t.userid = u.id
           WHERE (t.visibility = $3)
           OR  
	(
		t.userid IN  ( SELECT followee_id FROM followers WHERE follower_id = $4) 
		AND 
		t.visibility IN ('public' ,'followers')
	) OR (
		t.userid = $4 AND t.visibility  IN ('public','followers')
	)
	   ORDER BY t.created_at DESC
           LIMIT $1 OFFSET $2 `

	err := db.Select(&tweets, query, pageSize, offset, vis, userId)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}

	return tweets, nil
}
