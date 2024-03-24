package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func UpdatePasswordPayload(r *http.Request) (*models.UpdatePasswordPayload, error) {
	params := &models.UpdatePasswordPayload{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
