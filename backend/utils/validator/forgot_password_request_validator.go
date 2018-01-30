package validator

import "net/http"

//ForgotPasswordRequestData ..
type ForgotPasswordRequestData struct{
	Email string
}

//Validate ..
func (d *ForgotPasswordRequestData) Validate(r *http.Request) error {
	var err error
	err = ValidateEmail(d.Email)
	if err != nil {
		return err
	}
	err = ValidateEmailExists(d.Email)
	if err != nil {
		return err
	}
	return nil
}