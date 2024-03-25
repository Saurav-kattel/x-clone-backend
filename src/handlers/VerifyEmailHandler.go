package handlers

import (
	"database/sql"
	"log"
	"net/http"
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

		_, otpErr := user.GetOtp(db, &uid)
		if otpErr != nil {
			if otpErr != sql.ErrNoRows {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res:    models.Message{Message: otpErr.Error()},
				})
				return
			}
		} else {
			delErr := user.DeleteOtp(db, &uid)
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
		otpId, err := user.CreateOtp(db, otp, uid)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		go func(otp, email string) {
			//sending otp from mail to use
			mailErr := user.SendMail(otp, userData.Email)
			if mailErr != nil {
				log.Fatal(mailErr.Error())
			}
		}(otp, userData.Email)

		// generating jwt for otp verification
		otpToken, err := encoder.CreateOtpJwt(otp, otpId.String())
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res:    models.Message{Message: err.Error()},
			})
			return
		}

		// generating cookies and saving to http response header
		cookie := http.Cookie{
			Name:    "otp_auth",
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
