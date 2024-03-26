package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
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

		//using context to retrive data
		userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
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

		go func() {

			subject := "Password Reset"

			message := fmt.Sprintf("From: X-clone %s\r\n"+
				"Subject: %s\r\n"+
				"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
				"<p>Your account password was updated, if it was not your doing please change your password</p>"+
				os.Getenv("EMAIL"), subject)

			err := user.SendMail(userData.Email, message)
			if err != nil {
				log.Printf("Error UpdateForgonPass: %s", err.Error())
			}
		}()
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "update password successfully",
		})

	}
}
