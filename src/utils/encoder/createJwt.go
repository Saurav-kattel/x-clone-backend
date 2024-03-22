package encoder

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(email, userId string) (string, error) {

	if userId == "" || email == "" {
		return "", fmt.Errorf("required data not found")
	}

	//validating secret key
	secret := []byte(os.Getenv("SECRET"))
	if len(secret) < 30 {
		return "", fmt.Errorf("invalid secret token")
	}

	//signing token with user data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
	})

	//signing token with secret token
	jwtToken, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return jwtToken, err
}
