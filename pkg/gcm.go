package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

type AesGCM struct {
}

func (crpytor AesGCM) Encrypt(password, salt, inputPath, outputPath string) error {
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
	defer cipherTextFile.Close()

	_, err = cipherTextFile.Write(ciphertext)
	if err != nil {
		return err
	}

	return nil
}

func (crpytor AesGCM) Decrypt(password, salt, inputPath, outputPath string) error {
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
	defer plaintextFile.Close()

	_, err = plaintextFile.Write(decryptedData)
	if err != nil {
		return err
	}

	return nil
}
