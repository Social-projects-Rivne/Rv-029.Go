package validator

import (
	"net/http"
	"github.com/gocql/gocql"
)

//ForgotPasswordRequestData ..
type ConfirmRegistrationRequestData struct{
	Token string
	UUID  gocql.UUID
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