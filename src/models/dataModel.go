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
	LastName  string    `db:"last_name"`
	FirstName string    `db:"first_name"`
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
	Id             string     `db:"id" json:"id"`
	Content        string     `db:"content" json:"content"`
	AuthorUsername string     `db:"author_username" json:"author_username"`
	Author         string     ` json:"author"`
	ImageId        *string    `db:"imageid" json:"imageId"`
	CreatedAt      *time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updated_at" json:"updatedAt"`
	UserId         string     `db:"userid" json:"userId"`
	Visibility     *string    `db:"visibility" json:"visibility"`
}

type AuthorImage struct {
	Image []byte `db:"byte" json:"image"`
}
type TweetImage struct {
	Image []byte `db:"byte" json:"image"`
}

type UserData struct {
	Id        string    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	ImageId   any       `json:"imageId" db:"image_id"` //profile image id
	CreatedAt time.Time `json:"createdAt"  db:"created_at"`
	Role      string    `json:"role" db:"role"`
	LastName  string    `json:"lastName" db:"last_name"`
	FirstName string    `json:"firstName" db:"first_name"`
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

type UserRepliedTweet struct {
	Tweets
}

type NotificationData struct {
	ID                 *string    `db:"id" json:"id"`
	RecipientID        *string    `db:"recipient_id" json:"recipient_id"`
	RecipientUsername  *string    `db:"recipient_username" json:"recipient_username"`
	RecipientEmail     *string    `db:"recipient_email" json:"recipient_email"`
	RecipientImageID   *string    `db:"recipient_image_id" json:"recipient_image_id"`
	RecipientCreatedAt *time.Time `db:"recipient_created_at" json:"recipient_created_at"`
	RecipientLastName  *string    `db:"recipient_last_name" json:"recipient_last_name"`
	RecipientFirstName *string    `db:"recipient_first_name" json:"recipient_first_name"`
	ReciverID          *string    `db:"reciver_id" json:"reciver_id"`
	ReciverUsername    *string    `db:"reciver_username" json:"reciver_username"`
	ReciverEmail       *string    `db:"reciver_email" json:"reciver_email"`
	ReciverImageID     *string    `db:"reciver_image_id" json:"reciver_image_id"`
	ReciverCreatedAt   time.Time  `db:"reciver_created_at" json:"reciver_created_at"`
	ReciverLastName    *string    `db:"reciver_last_name" json:"reciver_last_name"`
	ReciverFirstName   *string    `db:"reciver_first_name" json:"reciver_first_name"`
	Message            string     `db:"message" json:"message"`
	Status             string     `db:"status" json:"status"`
	CreatedAt          time.Time  `db:"created_at" json:"created_at"`
	Type               *string    `db:"type" json:"type"`
	UpdatedAt          time.Time  `db:"updated_at" json:"updated_at"`
	TweetId            *string    `db:"tweet_id" json:"tweet_id"`
}
