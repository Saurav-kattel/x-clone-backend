package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
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
		userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
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

		// comparing password against hash
		validatePassword := validator.HashValidator(userData.Password, data.Password)
		if !validatePassword {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res: models.Message{
					Message: "incorrect password, try again!",
				},
			})
			return
		}

		deletionErr := user.DeleteAccount(db, userData.Id)

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
			MaxAge:   -1,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "account deleted successfully",
		})
	}

}
