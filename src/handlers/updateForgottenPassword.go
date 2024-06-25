package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

func UpdateForgottenPasswordHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//checking for req.Method
		if r.Method != "PUT" {
			encoder.ResponseWriter(w, http.StatusMethodNotAllowed, models.ErrorResponse{
				Status: http.StatusMethodNotAllowed,
				Res: models.Message{
					Message: "invalid http method",
				},
			})
			return
		}

		//decoding data coming from clinet into struct
		data, payloadErr := decoder.UpdateForgottenPasswordPayloadDecoder(r)
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

		otpAuthToken := r.Header.Get("change_password_x_clone")
		if otpAuthToken == "" || otpAuthToken == "undefined" {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "Otp Authorization token not provided"},
			})
			return
		}

		jwtData, err := validator.ValidateForgetPasswordJwt(otpAuthToken)
		//updating user password
		updatePasswordErr := user.UpdatePassword(db, jwtData.UserId, encoder.HashPassword(data.NewPassword))
		if updatePasswordErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: updatePasswordErr.Error()},
			})
			return
		}

		//spawning new goroutine to send main

		go func(email string) {
			subject := "Password Reset"

			message := fmt.Sprintf("From: X-clone %s\r\n"+
				"Subject: %s\r\n"+
				"Content-Type: text/html; charset=UTF-8\r\n\r\nstring"+
				"<p>Your account password was updated. If it was not you,please change your password</p>",
				email, subject)

			err := user.SendMail(email, message)
			if err != nil {
				log.Printf("Error UpdateForgonPass: %s", err.Error())
			}
		}(jwtData.Email)

		// deleting otp of the user
		deleteOtpErr := user.DeleteOtp(db, jwtData.UserId)
		if deleteOtpErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: deleteOtpErr.Error()},
			})
			return
		}

		//signing new jwt token with login credentials
		loginJwtToken, err := encoder.CreateJwt(jwtData.Email, jwtData.UserId)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		//setting auth cookie
		token := http.Cookie{
			Name:     "auth_token_x_clone",
			Value:    loginJwtToken,
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(w, &token)

		//deleting otp auth cookie
		deleteCookie := &http.Cookie{
			Name:     "change_password_x_clone",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		}
		http.SetCookie(w, deleteCookie)

		//sending status ok
		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "password  updated successfully",
		})

	}

}
