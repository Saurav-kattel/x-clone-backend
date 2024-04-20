package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/encoder"
)

func TweetLikeHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
			})
			return
		}

		tweetId := r.URL.Query().Get("t_id")
		if tweetId == "" || tweetId == "undefined" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "tweet id not found",
				},
			})
			return
		}

		queryErr := tweets.HasUserLiked(db, userData.Id, tweetId)
		if queryErr == sql.ErrNoRows {
			log.Print("queryErr", queryErr.Error())
			err := tweets.HandleLike(db, tweetId, userData.Id)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}
		} else if queryErr == nil {
			err := tweets.HandleUnlike(db, tweetId, userData.Id)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}
		} else {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: queryErr.Error(),
				},
			})
			return
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "action successful",
		})
	}
}
