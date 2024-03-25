package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func VerifyEmailDecoder(r *http.Request) (*models.VerifyEmail, error) {
	params := &models.VerifyEmail{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
