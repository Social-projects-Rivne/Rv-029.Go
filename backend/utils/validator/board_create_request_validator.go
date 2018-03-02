package validator

import (
	"log"
	"net/http"
)

type BoardCreateRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"description"`
}

func (b *BoardCreateRequestData) Validate(r *http.Request) error {
	var err error

	err = b.ValidateRequired(b.Name)
	if err != nil {
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)
		return err
	}

	err = b.ValidateRequired(b.Desc)
	if err != nil {
		log.Printf("Error in utils/validator/board_create_request_validator.go error: %+v",err)
		return err
	}

	return nil
}
