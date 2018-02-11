package validator

import (
	"net/http"
)

type BoardRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (b *BoardRequestData) Validate(r *http.Request) error {
	// just for Interface implementation
	return nil
}
