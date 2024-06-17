package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func UpdateVisDecoder(r *http.Request) (*models.TweetVisPayload, error) {
	params := &models.TweetVisPayload{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
