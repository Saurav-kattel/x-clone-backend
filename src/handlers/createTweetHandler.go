package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/middleware"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/tweets"
	"x-clone.com/backend/src/utils/decoder"
	"x-clone.com/backend/src/utils/encoder"
)

//need to create a seperate handler to insert tweet image

func CreateTweetHandler(db *sqlx.DB) http.HandlerFunc {
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
		// parsing req.FormValue
		data, err := decoder.TweetPayloadDecoder(r)
		if err != nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "req payload: " + err.Error(),
				},
			})
			return
		}

		//parsing multipart data and setting up 10MB limit
		multipartErr := r.ParseMultipartForm(10 << 20)
		// Check if the request contains multipart data
		if multipartErr != nil && multipartErr != http.ErrNotMultipart {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "error parsing multipart form: " + multipartErr.Error(),
				},
			})
			return
		}

		if r.MultipartForm == nil {
			encoder.ResponseWriter(w, http.StatusBadRequest, models.ErrorResponse{
				Status: http.StatusBadRequest,
				Res: models.Message{
					Message: "multipart is nil",
				},
			})
			return
		}

		//reading mulipart data from request object
		file, _, formFileErr := r.FormFile("file")
		if formFileErr != nil {
			tweets.HandleTweetsWithoutImage(w, db, data, userData.Id)
		} else {
			tweets.HandleTweetWithImage(w, db, file, data, userData.Id)
			defer file.Close()

		}

		encoder.ResponseWriter(w, http.StatusOK, models.SuccessResponse{
			Status: http.StatusOK,
			Res:    "tweets created successfully",
		})
	}
}
