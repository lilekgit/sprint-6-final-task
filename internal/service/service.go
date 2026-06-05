package service

import (
	"strings"

	"sprint6finaltask/pkg/morse"
)

func isMorse(s string) bool {
	hasDotOrDash := false
	for _, r := range s {
		switch r {
		case '.', '-':
			hasDotOrDash = true
		case ' ':
			// допустимый символ
		default:
			return false
		}
	}
	return hasDotOrDash
}

func AutoConvert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", nil
	}

	if isMorse(input) {
		return morse.ToText(input), nil
	}
	return morse.ToMorse(input), nil
}
