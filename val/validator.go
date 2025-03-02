package val

import (
	"fmt"
	"net/mail"
)

func ValidateString(value string, minLength int, maxLength int) error {
	if len(value) < minLength || len(value) > maxLength {
		return fmt.Errorf("string length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 20); err != nil {
		return err
	}
	if !isAlphaNumeric(value) {
		return fmt.Errorf("username must be alphanumeric")
	}
	return nil
}

func ValidatePassword(value string) error {
	if err := ValidateString(value, 6, 20); err != nil {
		return err
	}
	return nil
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 6, 50); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 50); err != nil {
		return err
	}
	return nil
}

func isAlphaNumeric(value string) bool {
	for _, char := range value {
		if !((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}
