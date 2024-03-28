package tweets

import (
	"image"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

func HandleTweetWithImage(w http.ResponseWriter, db *sqlx.DB, file io.Reader, data *models.TweetsPayload, userId string) {
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

	tweetId, err := CreateTweets(db, data, userId)
	if err != nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res: models.Message{
				Message: err.Error(),
			},
		})
		return
	}

	imgId, err := InsertTweetImage(db, compressedImage, tweetId)

	if err != nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res: models.Message{
				Message: err.Error(),
			},
		})
		return
	}

	updateImgErr := InsertImageId(db, tweetId, imgId)
	if updateImgErr != nil {
		encoder.ResponseWriter(w, http.StatusInternalServerError, models.ErrorResponse{
			Status: http.StatusInternalServerError,
			Res: models.Message{
				Message: updateImgErr.Error(),
			},
		})
		return
	}
}
