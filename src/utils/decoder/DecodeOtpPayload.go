package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func DecodeOtpPayload(r *http.Request) (*models.OtpPayload, error) {
	params := &models.OtpPayload{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil

}
