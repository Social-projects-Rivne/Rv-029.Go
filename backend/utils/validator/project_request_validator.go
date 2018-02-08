package validator

import "net/http"

type ProjectRequestData struct {
	*baseValidator
	Name string
}

func (d *ProjectRequestData) Validate(r *http.Request) error {
	err := d.ValidateRequired(d.Name)
	if err != nil {
		return err
	}

	return nil
}