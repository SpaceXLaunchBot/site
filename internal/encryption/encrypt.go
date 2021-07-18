package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Lots of the code in this package is edited from https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

// Encrypt encrypts the given data with a given 32 byte key.
func Encrypt(key, data []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return []byte{}, err
	}

	// Galois/Counter Mode, very good performance. https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return []byte{}, err
	}

	// Random nonce size of gcm.NonceSize.
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return []byte{}, err
	}

	// Seal encrypts and authenticates plaintext, authenticates the additional data and appends the result to dst,
	// returning the updated slice. The nonce must be NonceSize() bytes long and unique for all time, for a given key.
	return gcm.Seal(nonce, nonce, data, nil), nil
}
