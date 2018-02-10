package validator

import "net/http"

type BoardRequestData struct {
	*baseValidator
	Name string `json:"name"`
	Desc string `json:"desc"`
	//ProjectId gocql.UUID // todo
}

func (b *BoardRequestData) Validate(r *http.Request) error {
	var err error

	err = b.ValidateRequired(b.Name)
	if err != nil {
		return err
	}

	err = b.ValidateRequired(b.Desc)
	if err != nil {
		return err
	}

	return nil
}
