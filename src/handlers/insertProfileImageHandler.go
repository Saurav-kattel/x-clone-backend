package handlers

import (
	"database/sql"
	"image"
	_ "image/png"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
	"x-clone.com/backend/src/utils/validator"
)

func InsertProfileHandler(db *sqlx.DB) http.HandlerFunc {
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

		//validating user
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

		userInfo, queryErr := user.GetUserByEmail(db, userData.Email)
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

		//parsing multipart data and setting up 10MB limit
		multipartParseErr := r.ParseMultipartForm(10 << 20)
		if multipartParseErr != nil {
			if multipartParseErr == http.ErrNotMultipart {
				encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
					Status: http.StatusBadRequest,
					Res: models.Message{
						Message: "request is not a multipart",
					},
				})
				return
			} else if multipartParseErr == io.EOF {
				encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
					Status: http.StatusBadRequest,
					Res: models.Message{
						Message: "unexpected end of multipart body",
					},
				})
				return
			} else {
				encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
					Status: http.StatusBadRequest,
					Res: models.Message{
						Message: "parsing error: " + multipartParseErr.Error(),
					},
				})
				return
			}
		}
		//reading mulipart data from request object
		file, data, formFileErr := r.FormFile("file")
		log.Print(file)
		log.Print(data.Filename)
		log.Print(data.Header)
		if formFileErr != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "reading form file error: " + formFileErr.Error(),
				},
			})
			return
		}
		defer file.Close()

		//decoding file into image.Image type
		orginalImage, _, err := image.Decode(file)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: "image deocoed err: " + err.Error(),
				},
			})
			return
		}

		//compressing image using jpeg compressing algrotim
		compressedImage, err := encoder.CompressImage(orginalImage)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		//parsing string into uuid
		parsedUserId, err := uuid.Parse(userData.UserId)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: err.Error(),
				},
			})
			return
		}

		var imageUuid *uuid.UUID
		//updating profile image if there exits previous image
		if userInfo.ImageId != "" {
			log.Print("updated")
			uid, err := user.UpdateProfileImage(db, parsedUserId, compressedImage)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}

			imageUuid = uid

		} else {
			//inserting image into db if there doesn't exist previous image
			uid, err := user.InsertProfileImage(db, parsedUserId, compressedImage)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}

			imageUuid = uid
		}

		insertImageIdErr := user.InsertImageId(db, imageUuid.String(), userData.UserId)
		if insertImageIdErr != nil {
			encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
				Status: http.StatusInternalServerError,
				Res: models.Message{
					Message: insertImageIdErr.Error(),
				},
			})
			return
		}
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "updated image successfully",
		})
	}
}
