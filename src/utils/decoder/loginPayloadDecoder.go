package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

// dumping request body to json

func LoginPayloadJsonDecoder(r *http.Request) (*models.LoginPayload, error) {
	params := &models.LoginPayload{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
