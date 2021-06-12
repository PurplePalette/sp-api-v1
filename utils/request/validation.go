package request

import "regexp"

// IsValidName validates the name is only number and alphabets.
func IsValidName(name string) bool {
	isAlpha := regexp.MustCompile(`^[0-9a-zA-Z]+$`).MatchString
	return isAlpha(name)
}
