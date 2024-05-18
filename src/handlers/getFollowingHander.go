package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func GetFollowingList(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//validating http method
		if r.Method != "GET" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		userId := r.URL.Query().Get("u_id")

		if userId != "" && userId != "undefined" {
			data, err := user.GetFollowingList(db, userId) // wasted 2 hours just because i called the wrong function

			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return

			}

			//sending status ok
			encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
				Status: http.StatusOK,
				Res:    data,
			})
			return
		} else {
			userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
			if !ok {
				encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
					Status: http.StatusUnauthorized,
					Res:    models.Message{Message: "User information not found"},
				})
				return
			}
			data, err := user.GetFollowingList(db, userData.Id)

			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return

			}

			//sending status ok
			encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
				Status: http.StatusOK,
				Res:    data,
			})
		}
	}
}
