package api

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/pkg/errors"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt returns encrypted given data.
func Encrypt(data []byte, secret string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(createHash(secret)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, errors.WithStack(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt returns the data for given encrypted data.
func Decrypt(data []byte, secret string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(createHash(secret)))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return plaintext, nil
}
