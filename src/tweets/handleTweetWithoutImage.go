package tweets

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

func HandleTweetsWithoutImage(w http.ResponseWriter, db *sqlx.DB, data *models.TweetsPayload, userId string) {
	if db == nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res: models.Message{
				Message: "database connection is nil",
			},
		})
		return
	}

	if data == nil {
		encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
			Status: http.StatusBadRequest,
			Res: models.Message{
				Message: "tweets payload is nil",
			},
		})
		return
	}

	_, err := CreateTweets(db, data, userId)
	if err != nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res: models.Message{
				Message: err.Error(),
			},
		})
		return
	}
}
