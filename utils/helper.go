package utils

import (
	"regexp"
	"strings"
)

// Slugify returns a slugified version of the input string.
func Slugify(text string) string {
	// Convert the string to lowercase
	text = strings.ToLower(text)

	// Remove all non-alphanumeric characters except spaces
	re := regexp.MustCompile(`[^\w\s-]`)
	text = re.ReplaceAllString(text, "")

	// Replace spaces and underscores with hyphens
	text = strings.ReplaceAll(text, " ", "-")
	text = strings.ReplaceAll(text, "_", "-")

	// Replace multiple hyphens with a single hyphen
	re = regexp.MustCompile(`-+`)
	text = re.ReplaceAllString(text, "-")

	// Trim any leading or trailing hyphens
	text = strings.Trim(text, "-")

	return text
}
