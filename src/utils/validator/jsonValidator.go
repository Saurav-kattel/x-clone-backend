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

	default:
		return &models.ValidatorResponse{
			Field:   "payload",
			Message: "unsupported payload type",
		}
	}

	return nil
}
