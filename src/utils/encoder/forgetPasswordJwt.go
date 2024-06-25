package encoder

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func ForgetPasswordJwt(userId, otpId, email string) (string, error) {

	if otpId == "" || userId == "" || email == "" {
		return "", fmt.Errorf("required data not found")
	}

	//validating secret key
	secret := []byte(os.Getenv("SECRET"))
	if len(secret) < 30 {
		return "", fmt.Errorf("invalid secret token")
	}

	//signing token with otp and encoded data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"otpId":  otpId,
		"userId": userId,
		"email":  email,
	})

	//signing token with secret token
	jwtToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return jwtToken, err

}
