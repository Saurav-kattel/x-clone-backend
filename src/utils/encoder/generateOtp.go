package encoder

import (
	"math/rand"
	"strconv"
)

func GenerateOtp() string {

	otp := ""
	for i := 0; i < 4; i++ {
		randNum := rand.Intn(10)
		otp += strconv.Itoa(randNum)
	}
	return otp
}
