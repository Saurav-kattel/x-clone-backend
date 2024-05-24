package handlers

import (
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/encoder"
)

func GetTweetsHandler(db *sqlx.DB) http.HandlerFunc {
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
		username := r.URL.Query().Get("u_name")
		pageNumberStr := r.URL.Query().Get("n")
		pageSizeStr := r.URL.Query().Get("s")
		pageNumber, err := strconv.Atoi(pageNumberStr)
		tweetId := r.URL.Query().Get("t_id")

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

		if username != "" && username != "undefined" {
			tweets, err := tweets.GetUserPost(db, pageSize, pageNumber, username)
			if err != nil {
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
				Res:    tweets,
			})
			return
		} else if tweetId != "" && tweetId != "undefined" {
			tweets, err := tweets.GetTweetsById(db, tweetId)
			if err != nil {
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
				Res:    tweets,
			})
		} else {

			tweets, err := tweets.GetTweets(db, pageNumber, pageSize)
			if err != nil {
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
				Res:    tweets,
			})
		}
	}
}
