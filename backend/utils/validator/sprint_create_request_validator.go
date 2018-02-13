package validator

import (
	"net/http"
	"log"
)

type SprintCreateRequestData struct {
	*baseValidator
	Goal   string `json:"goal"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

func (s *SprintCreateRequestData) Validate(r *http.Request) error {
	var err error

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

	return validateBoardId(r)
}
