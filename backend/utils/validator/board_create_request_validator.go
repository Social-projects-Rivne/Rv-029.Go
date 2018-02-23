package validator

import (
	"log"
	"net/http"
)


type BoardCreateRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (b *BoardCreateRequestData) Validate(r *http.Request) error {
	var err error

	err = b.ValidateRequired(b.Name)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	err = b.ValidateRequired(b.Desc)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}
