package decoder

import (
	"encoding/json"
	"net/http"

	"x-clone.com/backend/src/models"
)

func CommentsPayloadDecoder(r *http.Request) (*models.Comment, error) {
	params := &models.Comment{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		return nil, err
	}

	return params, nil
}
