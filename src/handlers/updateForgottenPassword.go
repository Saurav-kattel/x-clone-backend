package handlers

import (
	"database/sql"
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

		//  validating otp auth token
		otpAuthToken := r.Header.Get("otp_auth_x_clone")
		if otpAuthToken == "" {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "Otp Authorization token not provided"},
			})
			return
		}

		// decoding necessary data from otp
		jwtData, err := validator.ValidateOtpJwt(otpAuthToken)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		// getting otp associated with user
		otpData, err := user.GetOtpWithUser(db, jwtData.UserId)
		if err != nil {
			if err == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res:    models.Message{Message: "otp not found"},
				})
			} else {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: err.Error()},
				})
			}
			return
		}

		// confirming  otp from req against database
		err, ok := validator.ConfirmOtp(data.Otp, otpData.Otp, otpData.ExpiresAt)

		if !ok {
			//checking if the error was or not the expired otp error
			if err.Error() != "expired otp" {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: "invalid opt, try again!"},
				})
				return
			}

			//deleting the expired otp from the database

			deletionErr := user.DeleteOtp(db, jwtData.UserId)
			if deletionErr != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: deletionErr.Error()},
				})
				return
			}

			//sending the error  for expired otp.
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: "expired otp. Please request a new one"},
			})
			return
		}

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
		}(otpData.Email)

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
		loginJwtToken, err := encoder.CreateJwt(otpData.Email, otpData.UserId)
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
			Name:     "otp_auth_x_clone",
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
