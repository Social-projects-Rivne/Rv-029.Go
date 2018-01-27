package validator

import "net/http"

type LoginRequestData struct {
	Email string
	Password string
}

func (d *LoginRequestData) Validate(r *http.Request) error {
	err := ValidateEmail(d.Email)
	if err != nil {
		return err
	}

	return nil
}