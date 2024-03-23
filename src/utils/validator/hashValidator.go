package validator

import (
	"crypto/sha256"
	"encoding/hex"
)

// function to hash password using sha256 algorithm

func HashValidator(hash, password string) bool {
	bytePass := []byte(password)
	passHash := sha256.Sum256(bytePass)
	hexHash := hex.EncodeToString(passHash[:])

	return hexHash == hash
}
