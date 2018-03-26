package validator

import (
	"net/http"
	"log"	
)

type UserProjectRequestData struct {
	*baseValidator
	Role string  `json:"role"`
	UserID string `json:"user"`
}

func (d *UserProjectRequestData) Validate(r *http.Request) error {
	err := d.ValidateRequired(d.UserID)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)

		return err
	}


	return nil
}
