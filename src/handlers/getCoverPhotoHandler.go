package handlers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func GetCoverPhotoHandler(db *sqlx.DB) http.HandlerFunc {
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

		if username == "" || username == "undefined" {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "username not found",
				},
			})
			return
		}

		userData, err := user.GetUserByUsername(db, username)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
				Status: http.StatusNotFound,
				Res:    models.Message{Message: "profile image not found"},
			})
			return
		}

		imageData, err := user.GetCoverImage(db, userData.Id)
		if err != nil {
			if err == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res:    models.Message{Message: "profile image not found"},
				})
				return
			} else {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
				return
			}
		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    imageData,
		})
	}
}
