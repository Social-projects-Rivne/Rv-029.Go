package validator

import (
	"log"
	"net/http"
)

type SprintCreateRequestData struct {
	*baseValidator
	Goal   string `json:"goal"`
	Desc   string `json:"desc"`
	Status string
}

func (s *SprintCreateRequestData) Validate(r *http.Request) error {
	var err error

	s.Status = "Todo"

	err = s.ValidateRequired(s.Goal)
	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)
		return err
	}

	err = s.ValidateRequired(s.Desc)
	if err != nil {
		log.Printf("Error in utils/validator/sprint_create_request_validator.go error: %+v",err)
		return err
	}

	return nil
}
