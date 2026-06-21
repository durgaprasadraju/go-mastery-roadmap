package solutions

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashPasswordSHA256 hashes password (use bcrypt/argon2 in production).
func HashPasswordSHA256(password string) string {
	h := sha256.Sum256([]byte(password))
	return hex.EncodeToString(h[:])
}

// ConstantTimeCompare compares strings in constant time (simplified).
func ConstantTimeCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := 0; i < len(a); i++ {
		result |= a[i] ^ b[i]
	}
	return result == 0
}