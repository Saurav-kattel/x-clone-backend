package handlers

import (
	"image"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/user"
	"x-clone.com/backend/src/utils/encoder"
)

func InsertCoverPictureHandler(db *sqlx.DB) http.HandlerFunc {
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

		userData, ok := r.Context().Value(middleware.UserContextKey).(*models.User)
		if !ok {
			encoder.ResponseWriter(w, http.StatusUnauthorized, models.ErrorResponse{
				Status: http.StatusUnauthorized,
				Res:    models.Message{Message: "User information not found"},
			})
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
		file, _, formFileErr := r.FormFile("file")
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

		data, _ := user.GetCoverImage(db, userData.Id)

		if data != nil {
			//updating  cover image if there exits previous image
			err := user.UpdateCoverImage(db, userData.Id, compressedImage)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}

		} else {
			//inserting image into db if there doesn't exist previous image
			err := user.CreateCoverImage(db, userData.Id, compressedImage)
			if err != nil {
				encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
					Status: http.StatusInternalServerError,
					Res: models.Message{
						Message: err.Error(),
					},
				})
				return
			}
		}
		encoder.ResponseWriter(w, 200, models.SuccessResponse{
			Status: 200,
			Res:    "updated image successfully",
		})
	}
}
