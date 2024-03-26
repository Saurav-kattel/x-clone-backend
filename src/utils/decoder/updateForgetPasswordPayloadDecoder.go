package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func UpdateForgottenPasswordPayloadDecoder(r *http.Request) (*models.UpdateForgottenPasswordPayload, error) {
	params := &models.UpdateForgottenPasswordPayload{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
