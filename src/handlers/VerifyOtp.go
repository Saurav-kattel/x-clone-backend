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

func VerifyOtp(db *sqlx.DB) http.HandlerFunc {

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

		//  validating otp auth token
		otpAuthToken := r.Header.Get("otp_auth_x_clone")
		if otpAuthToken == "" {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "Otp Authorization token not provided"},
			})
			return
		}

		// decoding necessary data from otpauthtoken
		jwtData, err := validator.ValidateOtpJwt(otpAuthToken)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}
		// decoding req.body
		reqPayload, err := decoder.DecodeOtpPayload(r)
		if err != nil {

			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		if jwtData.Otp != reqPayload.Otp {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: "otp didnot match"},
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
		err, ok := validator.ConfirmOtp(reqPayload.Otp, otpData.Otp, otpData.ExpiresAt)
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

		jwt, err := encoder.ForgetPasswordJwt(otpData.UserId, otpData.Id, otpData.Email)

		cookie := &http.Cookie{
			Name:     "change_password_x_clone",
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(time.Minute * 10),
			Value:    jwt,
		}
		http.SetCookie(w, cookie)

		deleteCookie := &http.Cookie{
			Name:     "otp_auth_x_clone",
			Value:    "",
			Path:     "/",
			Expires:  time.Unix(0, 0),
			HttpOnly: true,
		}
		http.SetCookie(w, deleteCookie)

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "success",
		})
	}
}
