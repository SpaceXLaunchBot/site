package encryption

import "crypto/rand"

// GenerateKey generates a key for use with AES GCM.
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}
