package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func RegisterPayloadJsonDecoder(r *http.Request) (*models.RegisterPayload, error) {
	params := &models.RegisterPayload{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
