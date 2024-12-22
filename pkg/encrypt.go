package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Encrypt(password, salt, inputPath, outputPath string) error {
	secretKey := CreateSecretKey(password, salt)

	plaintext, err := os.ReadFile(inputPath)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("error setting gcm mode: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("error generating the nonce: %v", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	cipherTextFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = cipherTextFile.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}
