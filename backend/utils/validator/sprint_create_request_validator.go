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
		log.Printf(err.Error())
		return err
	}

	err = s.ValidateRequired(s.Desc)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}
