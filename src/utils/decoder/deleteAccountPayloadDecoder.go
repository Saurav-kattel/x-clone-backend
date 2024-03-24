package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

// dumping request body to json

func DeleteAccountPayload(r *http.Request) (*models.DeleteAccountPayload, error) {
	params := &models.DeleteAccountPayload{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
