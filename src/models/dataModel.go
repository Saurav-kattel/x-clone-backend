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
	User_id  string `db:"user_id" json:"userId"`
	Username string `db:"username" json:"username"`
}

type Follow struct {
	FolloweeId string `json:"followeeId"`
}

type CommentData struct {
	Id              string    `db:"id" json:"id"`
	Comment         string    `db:"comment" json:"comment"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
	UserId          string    `db:"user_id" json:"userId"`
	Username        string    `db:"username" json:"username"`
	TweetId         string    `db:"tweet_id" json:"tweet_id"`
	ParentCommentId *string   `db:"parent_comment_id" json:"parentTweetId"`
}

type FollowerList struct {
	Username string `db:"username" json:"username"`
	UserId   string `db:"user_id" json:"user_id"`
	Id       string `db:"id" json:"id"`
}

type ReplyData struct {
	Id                  string  `db:"id" json:"id"`
	Reply               string  `json:"reply" db:"reply"`
	TweetId             string  `json:"tweet_id" db:"tweet_id"`
	RepliedTo           string  `json:"replied_to" db:"replied_to"`
	RepliedFrom         string  `json:"replied_from" db:"replied_from"`
	RepliedToUsername   string  `json:"replied_to_username" db:"replied_to_username"`
	RepliedFromUsername string  `json:"replied_from_username" db:"replied_from_username"`
	CreatedAt           string  `json:"created_at" db:"created_at"`
	CommentId           string  `json:"comment_id" db:"comment_id"`
	Parent_id           *string `json:"parent_id" db:"parent_id"`
}
