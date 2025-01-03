package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// SanitizeLogin sanitizes the login by replacing the middle part with the specified character.
func SanitizeLogin(login string) string {
	// Check if the login is an email or phone number
	if isEmail(login) {
		return sanitizeEmail(login)
	} else if isPhoneNumber(login) {
		return sanitizePhoneNumber(login)
	}

	asterisks := strings.Repeat("*", len(login)-1)

	return fmt.Sprintf("%c%s", login[0], asterisks)
}

// isEmail checks if the input string is a valid email address.
func isEmail(email string) bool {
	// Basic email format validation using regular expression.
	// Note: This is a simple regex and may not cover all edge cases of valid email addresses.
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// sanitizeEmail replaces the middle part of the email with '*' for privacy.
func sanitizeEmail(email string) string {
	const replacedCharLength = 4

	localPart, domain := splitEmail(email)
	hiddenPart := string('*') + strings.Repeat("*", replacedCharLength) + string('*')

	return localPart[:1] + hiddenPart + localPart[len(localPart)-1:] + "@" + domain
}

// splitEmail splits an email into local part and domain.
func splitEmail(email string) (string, string) {
	var (
		localPart string
		domain    string
	)

	atIndex := strings.LastIndex(email, "@")
	if atIndex != -1 {
		localPart = email[:atIndex]
		domain = email[atIndex+1:]
	}

	return localPart, domain
}

// isPhoneNumber checks if the input string is a valid phone number.
func isPhoneNumber(phoneNumber string) bool {
	// Basic phone number format validation using regular expression
	// Note: This is a simple regex and may not cover all edge cases of valid phone numbers.
	phoneNumberRegex := regexp.MustCompile(`^\+?[0-9]*$`)
	return phoneNumberRegex.MatchString(phoneNumber)
}

// sanitizePhoneNumber replaces the middle part of the phone number with '*' for privacy.
func sanitizePhoneNumber(phoneNumber string) string {
	// Get the country code (if present) by checking for a leading '+'
	countryCode := ""
	countryCodeLength := 4

	if len(phoneNumber) < countryCodeLength {
		return phoneNumber
	}

	if phoneNumber[0] == '+' {
		countryCode = phoneNumber[:countryCodeLength]
		phoneNumber = phoneNumber[countryCodeLength:]
	}
	// Calculate the number of digits to keep at the end
	numDigitsToKeep := 4
	// Replace the middle digits with '*'
	hiddenPart := string('*')
	if len(phoneNumber) > numDigitsToKeep {
		hiddenPart = strings.Repeat("*", len(phoneNumber)-numDigitsToKeep-countryCodeLength)
	}
	// Return the sanitized phone number
	return countryCode + hiddenPart + phoneNumber[len(phoneNumber)-numDigitsToKeep:]
}
