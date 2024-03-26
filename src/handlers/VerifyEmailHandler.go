package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

func VerifyEmailHandler(db *sqlx.DB) http.HandlerFunc {
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

		// decoding req payload into struct
		data, payloadErr := decoder.VerifyEmailDecoder(r)
		if payloadErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: payloadErr.Error(),
				},
			})
			return
		}

		// validating req payload
		validationErr := validator.ValidatePayload(data)
		if validationErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.SuccessResponse{
				Status: http.StatusBadRequest,
				Res:    validationErr,
			})
			return
		}

		// fetching user data by payload email
		userData, queryErr := user.GetUserByEmail(db, data.Email)
		if queryErr != nil {
			if queryErr == sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusNotFound, models.ErrorResponse{
					Status: http.StatusNotFound,
					Res:    models.Message{Message: "User not found"},
				})
			} else {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: queryErr.Error()},
				})
			}
			return
		}
		uid, err := uuid.Parse(userData.Id)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		_, otpErr := user.GetOtpWithUser(db, userData.Id)
		if otpErr != nil {
			if otpErr != sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: otpErr.Error()},
				})
				return
			}
		} else {
			delErr := user.DeleteOtp(db, userData.Id)
			if delErr != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: delErr.Error()},
				})
				return
			}
		}

		// generetaing otp
		otp := encoder.GenerateOtp()
		// saving otp  into db and retriving data
		otpCreateErr := user.CreateOtp(db, otp, uid)
		if otpCreateErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: otpCreateErr.Error()},
			})
			return
		}

		//sending message in another go routine for faster api response
		go func(otp, email string) {

			subject := "Your OTP for Password Reset"

			message := fmt.Sprintf("From: X-clone %s\r\n"+
				"Subject: %s\r\n"+
				"Content-Type: text/html; charset=UTF-8\r\n\r\n"+
				"<h1>OTP</h1>"+
				"<p>Your OTP for changing the password is: <em><strong>%s</strong></em>.</p>"+
				"<p>Please do not share this OTP with others.</p>", os.Getenv("EMAIL"), subject, otp)

			//sending otp from mail to use
			mailErr := user.SendMail(userData.Email, message)
			if mailErr != nil {
				log.Fatal(mailErr.Error())
			}
		}(otp, userData.Email)

		// generating jwt for otp verification
		otpToken, err := encoder.CreateOtpJwt(otp, userData.Id)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		// generating cookies and saving to http response header
		cookie := http.Cookie{
			Name:    "otp_auth_x_clone",
			Value:   otpToken,
			Path:    "/",
			Expires: time.Now().Add(time.Minute * 10),
		}
		http.SetCookie(w, &cookie)

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "email  verified successfully",
		})
	}
}
