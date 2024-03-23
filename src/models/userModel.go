package models

import "time"

// user model or schema
type User struct {
	Id        string    `db:"id"`
	Username  string    `db:"username"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	ImageId   any       `db:"image_id"` //profile image id
	CreatedAt time.Time `db:"created_at"`
	Role      string    `db:"role"`
}
