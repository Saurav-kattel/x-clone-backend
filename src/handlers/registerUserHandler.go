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

func RegisterUserHandler(db *sqlx.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//checking for the request method
		if r.Method != "POST" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		//decoding req body to json
		data, err := decoder.RegisterPayloadJsonDecoder(r)
		if err != nil {
			encoder.ResponseWriter(w, 400, models.ErrorResponse{
				Status: 400,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		//validating request data
		validationErr := validator.ValidatePayload(data)
		if validationErr != nil {
			encoder.ResponseWriter(w, 400, models.SuccessResponse{
				Status: 400,
				Res:    validationErr,
			})
			return
		}

		//finding corrosponding user by email
		_, queryErr := user.GetUserByEmail(db, data.Email)
		if queryErr != nil && queryErr != sql.ErrNoRows {
			encoder.ResponseWriter(w, 500, models.ErrorResponse{
				Status: 500,
				Res: models.Message{
					Message: queryErr.Error(),
				},
			})
			return
		}
		// retuting if user is already available
		if queryErr == nil {
			encoder.ResponseWriter(w, 401, models.ErrorResponse{
				Status: 401,
				Res: models.Message{
					Message: "user already exists",
				},
			})
			return
		}

		_, userByUsernameErr := user.GetUserByUsername(db, data.Username)
		if userByUsernameErr != nil && userByUsernameErr != sql.ErrNoRows {
			encoder.ResponseWriter(w, 500, models.ErrorResponse{
				Status: 500,
				Res: models.Message{
					Message: userByUsernameErr.Error(),
				},
			})
			return
		}

		if userByUsernameErr == nil {
			encoder.ResponseWriter(w, 400, models.ErrorResponse{
				Status: 400,
				Res: models.Message{
					Message: "user with this username already exists",
				},
			})
			return

		}
		// creating a new user
		userId, createError := user.CreateUser(db, data)
		if createError != nil {
			encoder.ResponseWriter(w, 500, models.ErrorResponse{
				Status: 500,
				Res: models.Message{
					Message: createError.Error(),
				},
			})
			return
		}

		//signing jwt with user data
		token, tokenErr := encoder.CreateJwt(data.Email, *userId)
		if tokenErr != nil {
			encoder.ResponseWriter(
				w,
				http.StatusInternalServerError,
				models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: tokenErr.Error(),
					},
				},
			)
			return

		}

		//creating and appending cookie  with response header
		cookie := &http.Cookie{
			Name:     "auth_token_x_clone",
			HttpOnly: true,
			Path:     "/",
			Value:    token,
		}
		http.SetCookie(w, cookie)

		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "signed up successfully",
		})

	}

}
