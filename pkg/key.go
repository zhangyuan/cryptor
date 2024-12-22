package pkg

import (
	"crypto/sha256"

	"golang.org/x/crypto/pbkdf2"
)

func CreateSecretKey(password, salt string) []byte {
	secretKey := pbkdf2.Key([]byte(password), []byte(salt), 10, 32, sha256.New)
	return secretKey
}
