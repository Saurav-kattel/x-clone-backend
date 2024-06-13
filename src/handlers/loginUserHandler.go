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

// login handler function
func LoginUserHandler(db *sqlx.DB) http.HandlerFunc {

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

		data, err := decoder.LoginPayloadJsonDecoder(r)

		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		// validating user submitted data
		validationError := validator.ValidatePayload(data)
		if validationError != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.SuccessResponse{
				Status: http.StatusBadRequest,
				Res:    validationError,
			})
			return
		}

		//searching user with corresponding email
		user, queryErr := user.GetUserByEmail(db, data.Email)
		if queryErr != nil {
			if queryErr == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res: models.Message{
						Message: "user not found",
					},
				})
			} else {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: queryErr.Error(),
					},
				})
			}
			return
		}

		// validating user password with the saved hash
		validatePassword := validator.HashValidator(user.Password, data.Password)

		if !validatePassword {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res: models.Message{
					Message: "incorrect password, try again!",
				},
			})
			return
		}

		//siging jwt token with user data
		token, tokenErr := encoder.CreateJwt(user.Email, user.Id)

		if tokenErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: tokenErr.Error(),
				},
			})
		}

		//setting up cookies
		cookie := &http.Cookie{
			Name:     "auth_token_x_clone",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(time.Hour * 10),
			Value:    token,
		}
		http.SetCookie(w, cookie)

		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "logged in successfully",
		})

	}

}
