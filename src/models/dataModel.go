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
	Id        string     `db:"id" json:"id"`
	Content   string     `db:"content" json:"content"`
	Author    string     ` json:"author"`
	ImageId   *string    `db:"imageid" json:"imageId"`
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `db:"updated_at" json:"updatedAt"`
	UserId    string     `db:"userid" json:"userId"`
}

type AuthorImage struct {
	Image []byte `db:"byte" json:"image"`
}
type TweetImage struct {
	Image []byte `db:"byte" json:"image"`
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

type LikedUsers struct {
	Id       string `db:"id" json:"like_id"`
	User_id  string `db:"userid" json:"userId"`
	Username string `db:"username" json:"username"`
}
