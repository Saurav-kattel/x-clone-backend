package validator

import (
	"errors"

	"x-clone.com/backend/src/models"
)

func ValidateVisibility(data *models.TweetsPayload) error {
	if (data.Visibility != "public") && (data.Visibility != "private") && (data.Visibility != "followers") {
		return errors.New("unknown visibility parameters")
	}
	return nil
}
