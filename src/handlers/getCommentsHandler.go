package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/encoder"
)

func GetCommentsHandler(db *sqlx.DB) http.HandlerFunc {
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

		pageNumberStr := r.URL.Query().Get("n")
		pageSizeStr := r.URL.Query().Get("s")

		pageNumber, err := strconv.Atoi(pageNumberStr)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}
		if tweetId == "" || tweetId == "undefined" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "tweet id not found",
				},
			})
			return
		}
		comments, err := tweets.GetComments(db, tweetId, pageNumber, pageSize)
		if err != nil {
			if err == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res: models.Message{
						Message: "no comments",
					},
				})
				return
			}
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    comments,
		})

	}
}
