package encryption

import (
	"crypto/aes"
	"crypto/cipher"
)

// Decrypt decrypts the given data with a given 32 byte key.
func Decrypt(key, data []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return []byte{}, err
	}

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
