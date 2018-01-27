package validator

import (
	"regexp"
	"errors"
	"fmt"
	"net/http"
)

type InputValidation interface {
	Validate(r *http.Request) error
}

func validateMaxLenght(value string, max int) error {
	if len(value) > max {
		return errors.New(fmt.Sprint("Length should be less than %v chars", max))
	}

	return nil
}

func ValidateMinLenght(value string, min int) error {
	if len(value) < min {
		return errors.New(fmt.Sprint("Length should be more than %v chars", min))
	}

	return nil
}

func ValidateRequired (value interface{}) error {
	if value == nil {
		return errors.New("Value cannot be empty")
	}

	return nil
}

func ValidateEmail (email string) error {
	emailRegexp := regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)

	if !emailRegexp.MatchString(email) {
		return errors.New("Invalid Email")
	}

	return nil
}

