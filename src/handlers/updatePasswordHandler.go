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

func UpdatePasswordHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//validating req.method
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
		data, payloadErr := decoder.UpdatePasswordPayload(r)
		if payloadErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: payloadErr.Error(),
				},
			})
			return
		}

		//validating user data
		validationErr := validator.ValidatePayload(data)
		if validationErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.SuccessResponse{
				Status: http.StatusBadRequest,
				Res:    validationErr,
			})
			return
		}

		// getting user by email to validate or confirm password
		userInfo, queryErr := user.GetUserByEmail(db, userData.Email)
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

		//checking user password
		validatePassword := validator.HashValidator(userInfo.Password, data.OldPassword)
		if !validatePassword {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res: models.Message{
					Message: "incorrect password, try again!",
				},
			})
			return
		}

		//updating user password
		updateErr := user.UpdatePassword(db, userInfo.Id, encoder.HashPassword(data.NewPassword))
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
			Res:    "update password successfully",
		})

	}
}
