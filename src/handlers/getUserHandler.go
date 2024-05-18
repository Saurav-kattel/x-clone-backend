package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func GetUserHandler(db *sqlx.DB) http.HandlerFunc {
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

		if username == "" {
			userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
			if !ok {
				encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
					Status: http.StatusUnauthorized,
					Res:    models.Message{Message: "User information not found"},
				})
				return
			}

			encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
				Status: http.StatusOK,
				Res: models.UserData{
					Id:        userData.Id,
					Username:  userData.Username,
					Email:     userData.Email,
					Role:      userData.Role,
					ImageId:   userData.ImageId,
					CreatedAt: userData.CreatedAt,
					LastName:  userData.LastName,
					FirstName: userData.FirstName,
				},
			})
		} else {

			userData, err := user.GetUserByUsername(db, username)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError,
					models.ErrorResponse{
						Status: http.StatusInternalServerError,
						Res: models.Message{
							Message: err.Error(),
						},
					},
				)
				return
			}

			encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
				Status: http.StatusOK,
				Res: models.UserData{
					Id:        userData.Id,
					Username:  userData.Username,
					Email:     userData.Email,
					Role:      userData.Role,
					ImageId:   userData.ImageId,
					CreatedAt: userData.CreatedAt,
					LastName:  userData.LastName,
					FirstName: userData.FirstName,
				},
			})

		}
	}
}
