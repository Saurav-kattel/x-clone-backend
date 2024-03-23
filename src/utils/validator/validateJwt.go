package validator

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type DecodedJwtVal struct {
	Email  string
	UserId string
}

func ValidateJwt(jwtToken string) (*DecodedJwtVal, error) {
	tokenString := "8d73bf5184f48e9a1eedc1e5b215cd"
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
		email, _ := claims["email"].(string)
		userId, _ := claims["userId"].(string)

		return &DecodedJwtVal{
			Email:  email,
			UserId: userId,
		}, nil
	}

	return nil, fmt.Errorf("invlaid token")
}
