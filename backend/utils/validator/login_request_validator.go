package validator

import "net/http"

type LoginRequestData struct {
	*baseValidator
	Email    string
	Password string
}

func (d *LoginRequestData) Validate(r *http.Request) error {
	err := d.ValidateEmail(d.Email)
	if err != nil {
		return err
	}

	return nil
}
