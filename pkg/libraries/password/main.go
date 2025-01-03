package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash of the given password.
// It returns the hashed password or an error if the hashing process fails.
func HashPassword(password string) (string, error) {
	// Generate bcrypt hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// CheckPasswordHash compares a password with its bcrypt hash.
// It returns true if the password matches the hash, or false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
