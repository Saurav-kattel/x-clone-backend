package tweets

import (
	"fmt"
	"image"
	"io"
	"net/http"

	"github.com/jmoiron/sqlx"
	"x-clone.com/backend/src/models"
	"x-clone.com/backend/src/utils/encoder"
)

func HandleTweetWithImage(w http.ResponseWriter, db *sqlx.DB, file io.Reader, data *models.TweetsPayload, userId string) error {
	orginalImage, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("image decoding error : %+v", err.Error())
	}

	//compressing image using jpeg compressing algrotim
	compressedImage, err := encoder.CompressImage(orginalImage)
	if err != nil {
		return err
	}

	tweetId, err := CreateTweets(db, data, userId)
	if err != nil {
		return err
	}

	imgId, err := InsertTweetImage(db, compressedImage, tweetId)

	if err != nil {
		return err
	}

	updateImgErr := InsertImageId(db, tweetId, imgId)
	if updateImgErr != nil {

		return updateImgErr
	}

	return nil
}
