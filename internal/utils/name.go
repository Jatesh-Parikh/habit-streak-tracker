package utils

import (
	"errors"
	"strings"
	"unicode"
)

func ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if len(name) == 0 {
		return errors.New("Name cannot be empty")
	}

	if len(name) < 2 {
		return errors.New("Name must be atleast 2 characters long")
	}

	if len(name) > 100 {
		return errors.New("Name cannot exceed 100 characters")
	}

	for _, char := range name {
		if unicode.IsLetter(char) {
			continue
		}

		if unicode.IsDigit(char) {
			continue
		}

		if char == '.' {
			continue
		}

		if char == ' ' {
			continue
		}

		if char == '-' {
			continue
		}

		if char == '\'' {
			continue
		}

		return errors.New("Name can only contain letters, numbers, periods, spaces, hyphens, apostrophes")
	}

	return nil
}
