package handlers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/encoder"
)

func LikesCountHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		tweetId := r.URL.Query().Get("t_id")

		if tweetId == "" || tweetId == "undefined" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "invalid tweet id, got " + tweetId,
				},
			})
			return
		}
		count, err := tweets.GetLikesCount(db, tweetId)

		if err == sql.ErrNoRows {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.SuccessResponse{
				Status: http.StatusBadRequest,
				Res:    0,
			})
			return
		}
		if err != nil && err != sql.ErrNoRows {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
			})
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    count,
		})
	}
}
