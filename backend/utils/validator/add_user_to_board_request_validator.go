package validator

import (
	"net/http"
	"log"	
)

type UserBoardRequestData struct {
	*baseValidator
	Email string  `json:"email"`
	UserId string `json:"user_id"`
}

func (d *UserBoardRequestData) Validate(r *http.Request) error {
	err := d.ValidateRequired(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/project_request_validator.go error: %+v", err)		
		return err
	}

	err = d.ValidateRequired(d.UserId)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)

		return err
	}


	return nil
}
