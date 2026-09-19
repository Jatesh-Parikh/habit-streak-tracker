package utils

import (
	"errors"
	"unicode"
	// "golang.org/x/crypto/bcrypt"
)

func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("Password must be atleast 8 characters long")
	}

	var (
		hasUpper  bool
		hasLower  bool
		hasNumber bool
		hasSymbol bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true

		case unicode.IsLower(char):
			hasLower = true

		case unicode.IsDigit(char):
			hasNumber = true

		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		}
	}

	if !hasUpper {
		return errors.New("Password must contain atleast one uppercase letter")
	}

	if !hasLower {
		return errors.New("Password must contain atleast one lowercase letter")
	}

	if !hasNumber {
		return errors.New("Password must contain atleast one number")
	}

	if !hasSymbol {
		return errors.New("Password must contain atleast one symbol")
	}

	return nil
}
