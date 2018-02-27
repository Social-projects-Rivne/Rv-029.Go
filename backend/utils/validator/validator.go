package validator

import (
	"errors"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"net/http"
	"regexp"
)

//go:generate mockgen -destination=../../mocks/mock_validator.go -package=mocks github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator InputValidation

type InputValidation interface {
	Validate(r *http.Request) error
}

type baseValidator struct{}

func (v *baseValidator) ValidateMaxLenght(value string, max int) error {
	if len(value) > max {
		return errors.New(fmt.Sprintf("Length should be less than %v chars", max))
	}

	return nil
}

func (v *baseValidator) ValidateMinLenght(value string, min int) error {
	if len(value) < min {
		return errors.New(fmt.Sprintf("length should be more than %v chars", min))
	}

	return nil
}

func (v *baseValidator) ValidateRequired(value interface{}) error {
	if value == nil {
		return errors.New("value is not set")
	} else {
		switch value.(type) {
		case string:
			if value == "" {
				return errors.New("value cannot be empty string")
			}
			//TODO: add other types
		}
	}

	return nil
}

func (v *baseValidator) ValidateEmail(email string) error {
	emailRegexp := regexp.MustCompile(`(?:[a-z0-9!#$%&'*+/=?^_{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])`)

	if !emailRegexp.MatchString(email) {
		return errors.New("Invalid Email")
	}

	return nil
}

func (v *baseValidator) ValidateEmailUnique(email string) error {
	user := models.User{}
	user.Email = email
	models.UserDB.FindByEmail(&user)
	if user.Email != "" {
		return errors.New(fmt.Sprintf("User with %s email already exists", email))
	}

	return nil
}

func (v *baseValidator) ValidateEmailExists(email string) error {
	user := models.User{}
	user.Email = email
	models.UserDB.FindByEmail(&user)
	if user.Email == "" {
		return errors.New(fmt.Sprintf("User with %s email not exists", email))
	}

	return nil
}
