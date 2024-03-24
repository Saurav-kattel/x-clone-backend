package handlers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

func UpdateUsernameHandler(db *sqlx.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//validating http method
		if r.Method != "PUT" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		//checking for auth token in req.header and retrving it
		authToken := r.Header.Get("auth_token_x_clone")
		if authToken == "" {
			encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
				Status: http.StatusNotFound,
				Res: models.Message{
					Message: "auth token not found",
				},
			})
			return
		}

		// authenticating using jwt
		userData, err := validator.ValidateJwt(authToken)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		//decoding data coming from clinet into struct
		data, payloadErr := decoder.UpdateUsernamePayload(r)
		if payloadErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: payloadErr.Error(),
				},
			})
			return
		}

		// getting user by email to validate or confirm password
		_, queryErr := user.GetUserByEmail(db, userData.Email)
		if queryErr != nil {
			//sending error if user not found
			if queryErr == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res: models.Message{
						Message: "user not found",
					},
				})
			} else {
				//sending error for any other unexpected error
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: queryErr.Error(),
					},
				})
			}
			return
		}

		//updating user username
		updateErr := user.UpdateUsername(db, data.Username, userData.UserId)
		if updateErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: updateErr.Error(),
				},
			})
			return
		}

		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "update username successfully",
		})
	}
}
