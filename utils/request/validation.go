package request

import "regexp"

func IsValidName(name string) bool {
	isAlpha := regexp.MustCompile(`^[0-9a-zA-Z]+$`).MatchString
	return isAlpha(name)
}
