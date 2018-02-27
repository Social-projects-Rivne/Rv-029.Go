package validator

import (
	"net/http"
	"log"
)

//ForgotPasswordRequestData ..
type ResetPasswordRequestData struct {
	*baseValidator
	Email    string
	Password string
	Token    string
}

//Validate ..
func (d *ResetPasswordRequestData) Validate(r *http.Request) error {
	var err error

	err = d.ValidateEmail(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/reset_password_request_validator.go error: %+v",err)
		return err
	}

	err = d.ValidateEmailExists(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/reset_password_request_validator.go error: %+v",err)		
		return err
	}

	return nil
}
