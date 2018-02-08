package validator

import "net/http"

type ProjectRequestData struct {
	Name string
}

func (d *ProjectRequestData) Validate(r *http.Request) error {
	err := ValidateRequired(d.Name)
	if err != nil {
		return err
	}

	return nil
}