package validator

import "net/http"

//ForgotPasswordRequestData ..
type ForgotPasswordRequestData struct {
	*baseValidator
	Email string
}

//Validate ..
func (d *ForgotPasswordRequestData) Validate(r *http.Request) error {
	var err error

	err = d.ValidateEmail(d.Email)
	if err != nil {
		return err
	}

	err = d.ValidateEmailExists(d.Email)
	if err != nil {
		return err
	}

	return nil
}
