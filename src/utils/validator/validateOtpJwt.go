package validator

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// "otp":   otp,
// "otpId": otpId,

type OtpJwtData struct {
	Otp    string
	UserId string
}

func ValidateOtpJwt(jwtToken string) (*OtpJwtData, error) {
	tokenString := os.Getenv("SECRET")

	if len(tokenString) < 30 {
		return nil, fmt.Errorf("invalid secret token")
	}

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tokenString), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		otp, _ := claims["otp"].(string)
		userId, _ := claims["userId"].(string)

		return &OtpJwtData{
			Otp:    otp,
			UserId: userId,
		}, nil
	}

	return nil, fmt.Errorf("invlaid token")
}
