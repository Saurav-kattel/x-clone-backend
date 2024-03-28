package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func DeleteTweetPayloadDecoder(r *http.Request) (*models.DeleteTweetPayload, error) {
	params := &models.DeleteTweetPayload{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
