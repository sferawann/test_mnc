package utils

import (
	"errors"
	"regexp"
)

func ValidatePasswordMinLength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	return nil
}

func ValidateUsernameMinLength(username string) error {
	if len(username) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	return nil
}

func ValidateEmailFormat(email string) error {
	// Definisikan pola ekspresi reguler untuk format email yang valid
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Cocokkan email dengan pola ekspresi reguler
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return errors.New("invalid email format")
	}
	return nil
}
