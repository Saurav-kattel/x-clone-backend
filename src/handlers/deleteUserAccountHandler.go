package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

// endpoint handler for deleting user account
func DeleteUserAccountHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//validating http method
		if r.Method != "DELETE" {
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
		data, err := decoder.DeleteAccountPayload(r)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
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

		// comparing password against hash
		validatePassword := validator.HashValidator(userInfo.Password, data.Password)
		if !validatePassword {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res: models.Message{
					Message: "incorrect password, try again!",
				},
			})
			return
		}

		deletionErr := user.DeleteAccount(db, userData.UserId)

		if deletionErr != nil {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res: models.Message{
					Message: deletionErr.Error(),
				},
			})
			return
		}

		cookie := &http.Cookie{
			Name:     "auth_token_x_clone",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "account deleted successfully",
		})
	}

}
