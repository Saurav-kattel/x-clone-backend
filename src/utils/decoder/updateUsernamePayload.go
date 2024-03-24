package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func UpdateUsernamePayload(r *http.Request) (*models.UpdateUsernamePayload, error) {
	params := &models.UpdateUsernamePayload{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
