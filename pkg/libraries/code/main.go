package code

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomCode generates a random n-digit code.
// It seeds the random number generator with the current time.
func GenerateRandomCode(n int) string {
	// Set the possible characters that can be included in the code
	// In this example, we're using digits 0-9
	chars := "0123456789"

	code := make([]byte, n)
	for i := 0; i < n; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		code[i] = chars[index.Int64()]
	}

	return string(code)
}
