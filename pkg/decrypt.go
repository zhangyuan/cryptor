package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

func Decrypt(password, salt, inputPath, outputPath string) error {
	secretKey := CreateSecretKey(password, salt)

	ciphertext, err := os.ReadFile(inputPath)
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

	decryptedData, err := gcm.Open(nil, ciphertext[:gcm.NonceSize()], ciphertext[gcm.NonceSize():], nil)
	if err != nil {
		return err
	}

	plaintextFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = plaintextFile.Write(decryptedData)
	if err != nil {
		return err
	}

	return nil
}
