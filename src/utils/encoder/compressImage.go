package encoder

import (
	"bytes"
	"image"
	"image/jpeg"
)

func CompressImage(imageData image.Image) ([]byte, error) {
	const quality = 80

	var outputBuffer bytes.Buffer

	err := jpeg.Encode(&outputBuffer, imageData, &jpeg.Options{Quality: quality})
	if err != nil {
		return nil, err
	}

	return outputBuffer.Bytes(), nil

}
