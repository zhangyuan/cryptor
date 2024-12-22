package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type AesCfb struct {
}

func (cryptor AesCfb) Encrypt(password, salt, inputPath, outputPath string) error {
	secretKey := CreateSecretKey(password, salt)

	plaintextFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer plaintextFile.Close()

	cipherTextFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer cipherTextFile.Close()

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)

	if _, err = rand.Read(iv); err != nil {
		return err
	}

	if _, err := cipherTextFile.Write(iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	buffer := make([]byte, 4096)

	for {
		n, err := plaintextFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = cipherTextFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}

func (cryptor AesCfb) Decrypt(password, salt, inputPath, outputPath string) error {
	secretKey := CreateSecretKey(password, salt)

	ciphertextFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer ciphertextFile.Close()

	plaintextFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer plaintextFile.Close()

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return err
	}
	iv := make([]byte, aes.BlockSize)

	if _, err := io.ReadFull(ciphertextFile, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBDecrypter(block, iv)

	buffer := make([]byte, 4096)
	for {
		n, err := ciphertextFile.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		stream.XORKeyStream(buffer[:n], buffer[:n])
		_, err = plaintextFile.Write(buffer[:n])
		if err != nil {
			return err
		}
	}

	return nil
}
