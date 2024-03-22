package encoder

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	bytePassword := []byte(password)

	hash := sha256.Sum256(bytePassword)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}
