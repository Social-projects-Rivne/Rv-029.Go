package validator

import (
	"net/http"
	"log"
)

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
		log.Printf("Error in utils/validator/forgot_password_request_validator.go error: %+v",err)
		return err
	}

	err = d.ValidateEmailExists(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/forgot_password_request_validator.go error: %+v",err)		
		return err
	}

	return nil
}
