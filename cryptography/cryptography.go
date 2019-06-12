package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

const key = "shshk9119018hgaa"

func Encrypt(plaintext string) (string, error) {
	bytePlaintext := []byte(plaintext)
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encodedData := gcm.Seal(nonce, nonce, bytePlaintext, nil)
	return hex.EncodeToString(encodedData), nil
}

func Decrypt(ciphertext string) (string, error) {
	byteCiphertext, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, byteCiphertext := byteCiphertext[:nonceSize], byteCiphertext[nonceSize:]
	decodedData, err := gcm.Open(nil, nonce, byteCiphertext, nil)
	if err != nil {
		return "", err
	}
	return string(decodedData), nil
}
