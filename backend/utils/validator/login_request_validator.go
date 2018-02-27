package validator

import (
	"net/http"
	"log"
)

type LoginRequestData struct {
	*baseValidator
	Email    string
	Password string
}

func (d *LoginRequestData) Validate(r *http.Request) error {
	err := d.ValidateEmail(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/login_request_validator.go error: %+v",err)
		return err
	}

	return nil
}
