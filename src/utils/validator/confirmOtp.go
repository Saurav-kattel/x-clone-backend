package validator

import (
	"fmt"
	"time"
)

func ConfirmOtp(userOtp string, dbOtp string, expiresAt time.Time) (error, bool) {

	currentTime := time.Now().UTC()

	if userOtp != dbOtp {
		return fmt.Errorf("otp didnot matched"), false
	}

	if expiresAt.Before(currentTime) {
		return fmt.Errorf("exipred otp"), false
	}

	toleranceWindow := 5 * time.Second
	return nil, !expiresAt.Add(toleranceWindow).Before(currentTime)

}
