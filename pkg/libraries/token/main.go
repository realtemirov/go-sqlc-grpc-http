package token

import (
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const (
	Company = "company"
)

// Token structure with token value and expiry.
type Token struct {
	Type      string    `json:"type"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Function to generate a token string with an expiry time.
func GenerateTokenString(_type string, expiryDuration time.Duration) (string, error) {
	// Generate a new UUID as the token
	tokenValue := uuid.New().String()
	// Set expiry time
	expiryTime := time.Now().Add(expiryDuration)

	// Create a token struct
	token := Token{
		Token:     tokenValue,
		Type:      _type,
		ExpiresAt: expiryTime,
	}

	// Convert the token to JSON
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	// Encode the token JSON to Base64
	tokenString := base64.StdEncoding.EncodeToString(tokenJSON)

	return tokenString, nil
}

// Function to decode and check if a token string is expired.
func ParseToken(tokenString string) (*Token, error) {
	// Decode the Base64 encoded token
	tokenJSON, err := base64.StdEncoding.DecodeString(tokenString)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a Token struct
	var token Token
	err = json.Unmarshal(tokenJSON, &token)
	if err != nil {
		return nil, err
	}

	// Check if the token is expired
	return &token, nil
}
