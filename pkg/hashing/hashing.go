package hashing

import "crypto/sha256"

// Hashes the given byte slice using SHA256 and returns it
func HashSHA256(data []byte) []byte {
	result := sha256.Sum256(data)
	return result[:]
}
