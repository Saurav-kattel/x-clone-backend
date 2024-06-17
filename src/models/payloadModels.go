package models

type RegisterPayload struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	LastName        string `json:"lastName"`
	FirstName       string `json:"firstName"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DeleteAccountPayload struct {
	Password string `json:"password"`
}

type UpdateUsernamePayload struct {
	Username string `json:"username"`
}

type UpdatePasswordPayload struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type VerifyEmail struct {
	Email string `json:"email"`
}

type UpdateForgottenPasswordPayload struct {
	Otp             string `json:"otp"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type TweetsPayload struct {
	Content    string `json:"content"`
	Visibility string `json:"visibility"`
}

type DeleteTweetPayload struct {
	TweetId string `json:"tweetId"`
}

type Comment struct {
	Comment         string  `json:"comment"`
	TweetId         string  `json:"tweet_id"`
	ParentCommentId *string `json:"parent_comment_id"`
	RepliedTO       string  `json:"replied_to"`
	CommentId       string  `json:"comment_id"`
}
