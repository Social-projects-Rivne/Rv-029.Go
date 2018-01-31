package validator

import "net/http"

//ForgotPasswordRequestData ..
type ConfirmRegistrationRequestData struct{
	Token string
}

//Validate ..
func (d *ConfirmRegistrationRequestData) Validate(r *http.Request) error {
	var err error

	err = ValidateRequired(d.Token)
	if err != nil {
		return err
	}

	return nil
}