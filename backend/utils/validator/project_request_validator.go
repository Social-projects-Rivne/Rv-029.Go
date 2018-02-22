package validator

import (
	"net/http"
	"log"	
)

type ProjectRequestData struct {
	*baseValidator
	Name string
}

func (d *ProjectRequestData) Validate(r *http.Request) error {
	err := d.ValidateRequired(d.Name)
	if err != nil {
		log.Printf("Error in utils/validator/project_request_validator.go error: %+v", err)		
		return err
	}

	return nil
}