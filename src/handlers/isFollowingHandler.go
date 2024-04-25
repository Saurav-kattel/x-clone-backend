package handlers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func IsFollowing(db *sqlx.DB) http.HandlerFunc {
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

		followeeId := r.URL.Query().Get("f_id")

		if followeeId == "" || followeeId == "undefined" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "followee id is invalid",
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
		followingErr := user.IsFollowing(db, userData.Id, followeeId)

		if followingErr != nil && followingErr != sql.ErrNoRows {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: followingErr.Error()},
			})
			return
		}

		if followingErr == sql.ErrNoRows {
			encoder.ResponseWriter(w, 200, models.SuccessResponse{
				Status: 200,
				Res:    false,
			})
			return
		}
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    true,
		})

	}
}
