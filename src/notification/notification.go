package notification

import (
	"log"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
)

func CreateNotification(db *sqlx.DB, recipientId, reciverId, text, tweet_id *string, notType string) error {

	_, err := db.Exec(`INSERT INTO notifications(recipient_id, reciver_id, message, tweet_id,type) VALUES ($1, $2, $3,$4, $5)`, recipientId, reciverId, text, tweet_id, notType)
	log.Print(err)
	return err
}

func GetNotificationsByReceiverID(db *sqlx.DB, receiverId string) (*[]models.NotificationData, error) {
	var data []models.NotificationData
	query := `
        SELECT 
            notifications.id, 
            notifications.recipient_id, 
            notifications.reciver_id, 
            notifications.message, 
            notifications.status, 
            notifications.created_at, 
            notifications.updated_at,
            notifications.type,
            notifications.tweet_id,
            u1.id as "recipient_id",
            u1.username as "recipient_username",
            u1.email as "recipient_email",
            u1.image_id as "recipient_image_id",
            u1.created_at as "recipient_created_at",
            u1.last_name as "recipient_last_name",
            u1.first_name as "recipient_first_name",
            u2.id as "reciver_id",
            u2.username as "reciver_username",
            u2.email as "reciver_email",
            u2.image_id as "reciver_image_id",
            u2.created_at as "reciver_created_at",
            u2.last_name as "reciver_last_name",
            u2.first_name as "reciver_first_name"
        FROM 
            notifications 
        LEFT JOIN 
            users u1
        ON 
            notifications.recipient_id = u1.id
        JOIN
            users u2
        ON 
            notifications.reciver_id = u2.id
        WHERE 
            notifications.reciver_id = $1
        ORDER BY notifications.created_at DESC
	`
	err := db.Select(&data, query, receiverId)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func UpdateCommentStatus(db *sqlx.DB, notificationId string) error {
	_, err := db.Exec("UPDATE notifications SET status = 'read' WHERE id = $1", notificationId)
	return err
}
