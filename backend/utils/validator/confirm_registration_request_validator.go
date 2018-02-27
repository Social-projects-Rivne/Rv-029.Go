package validator

import (
	"github.com/gocql/gocql"
	"net/http"
	"log"
)

//ForgotPasswordRequestData ..
type ConfirmRegistrationRequestData struct {
	*baseValidator
	Token string
	UUID  gocql.UUID
}

//Validate ..
func (d *ConfirmRegistrationRequestData) Validate(r *http.Request) error {
	var err error

	err = d.ValidateRequired(d.Token)
	if err != nil {
		log.Printf("Error in utils/validator/confirm_registration_request_validator.go error: %+v",err)
		return err
	}

	return nil
}
