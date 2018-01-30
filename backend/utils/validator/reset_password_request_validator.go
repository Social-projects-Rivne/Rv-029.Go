package validator

import "net/http"

//ForgotPasswordRequestData ..
type ResetPasswordRequestData struct{
	Email string
	Password string
	Token string
}

//Validate ..
func (d *ResetPasswordRequestData) Validate(r *http.Request) error {
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