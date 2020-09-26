package services

import "strings"

//StringEmpty for checking string empty or not
func StringEmpty(str string) bool {
	if len(strings.TrimSpace(str)) == 0 {
		return true
	}

	return false
}

func IntegerEmpty(i int) bool {
	if i == 0 {
		return true
	}
	return false
}
