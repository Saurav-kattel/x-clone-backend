package validator

import (
	"regexp"

	"x-clone.com/backend/src/models"
)

func ValidatePayload(payload interface{}) *models.ValidatorResponse {

	// Type assertion to convert interface{} to the specific payload type
	switch p := payload.(type) {

	case *models.RegisterPayload: //for validating incoming new user register or signup payload

		if len(p.Username) < 3 {
			return &models.ValidatorResponse{
				Field:   "username",
				Message: "username  cannot be less the 3 characters long",
			}
		}

		if p.Email == "" {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "email field cannot be empty",
			}
		}

		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		regex := regexp.MustCompile(pattern)

		if !regex.MatchString(p.Email) {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "invalid email",
			}
		}

		if len(p.ConfirmPassword) <= 5 {
			return &models.ValidatorResponse{
				Field:   "confirm password",
				Message: "invalid  confirm password",
			}
		}

		if len(p.Password) <= 5 {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "invalid password",
			}
		}

		if p.ConfirmPassword != p.Password {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "password  did not match with re-entered password",
			}
		}

	case *models.LoginPayload:
		if p.Email == "" {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "email field cannot be empty",
			}
		}

		if len(p.Password) <= 5 {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "invalid password",
			}
		}

		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		regex := regexp.MustCompile(pattern)

		if !regex.MatchString(p.Email) {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "invalid email",
			}
		}
	case *models.UpdatePasswordPayload:

		if len(p.ConfirmPassword) <= 5 {
			return &models.ValidatorResponse{
				Field:   "confirm password",
				Message: "invalid  confirm password",
			}
		}

		if len(p.OldPassword) <= 5 {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "invalid password",
			}
		}

		if len(p.NewPassword) <= 5 {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "password too short",
			}
		}

		if p.ConfirmPassword != p.NewPassword {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "new password  did not match with confirm password",
			}
		}

	case *models.VerifyEmail:
		if p.Email == "" {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "email field cannot be empty",
			}
		}

		pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		regex := regexp.MustCompile(pattern)

		if !regex.MatchString(p.Email) {
			return &models.ValidatorResponse{
				Field:   "email",
				Message: "invalid email",
			}
		}

	case *models.UpdateForgottenPasswordPayload:
		if p.NewPassword == "" {
			return &models.ValidatorResponse{
				Field:   "new password",
				Message: "new password cannot be empty",
			}
		}

		if p.ConfirmPassword == "" {
			return &models.ValidatorResponse{
				Field:   "confirm password",
				Message: "confirm password cannot be empty",
			}
		}

		if p.ConfirmPassword != p.NewPassword {
			return &models.ValidatorResponse{
				Field:   "password",
				Message: "password  did not match with re-entered password",
			}
		}
	default:
		return &models.ValidatorResponse{
			Field:   "payload",
			Message: "unsupported payload type",
		}
	}

	return nil
}
