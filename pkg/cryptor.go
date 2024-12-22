package pkg

import (
	"fmt"
)

type Cryptor interface {
	Encrypt(password, salt, inputPath, outputPath string) error
	Decrypt(password, salt, inputPath, outputPath string) error
}

func GetCryptor(mode string) (Cryptor, error) {
	switch mode {
	case "gcm":
		return AesGCM{}, nil
	case "cfb":
		return AesCfb{}, nil
	default:
		return nil, fmt.Errorf("unsupported mode %v", mode)
	}
}

func Encrypt(mode, password, salt, inputFile, outputFile string) error {
	cryptor, err := GetCryptor(mode)
	if err != nil {
		return err
	}

	return cryptor.Encrypt(password, salt, inputFile, outputFile)
}

func Decrypt(mode, password, salt, inputFile, outputFile string) error {
	cryptor, err := GetCryptor(mode)
	if err != nil {
		return err
	}

	return cryptor.Decrypt(password, salt, inputFile, outputFile)
}
