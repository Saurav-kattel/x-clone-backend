package models

import (
	"time"

	"github.com/google/uuid"
)

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

type Otp struct {
	ID        uuid.UUID `db:"id"`
	Otp       string    `db:"otp"`
	UserId    string    `db:"userId"`
	CreatedAt time.Time `db:"created_at"`
	ExpiresAt time.Time `db:"expres_at"`
}

type OtpWithUser struct {
	Otp  Otp
	User User
}

type Tweets struct {
}

type UserData struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	ImageId   any       `json:"imageId"` //profile image id
	CreatedAt time.Time `json:"createdAt"`
	Role      string    `json:"role"`
}

type ProfileImage struct {
	Id    string `json:"id"`
	Image []byte `json:"image"`
}
