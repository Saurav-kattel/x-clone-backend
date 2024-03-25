package models

type RegisterPayload struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
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
