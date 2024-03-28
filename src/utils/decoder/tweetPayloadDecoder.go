package decoder

import (
	"encoding/json"
	"net/http"
	"strings"

	"x-clone.com/backend/src/models"
)

func TweetPayloadDecoder(r *http.Request) (*models.TweetsPayload, error) {
	params := &models.TweetsPayload{}
	jsonData := r.FormValue("data")
	reader := strings.NewReader(jsonData)
	err := json.NewDecoder(reader).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
