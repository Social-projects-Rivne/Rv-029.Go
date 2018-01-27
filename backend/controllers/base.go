package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
)

// decodeAndValidate - entrypoint for deserialization and validation
// of the submission input
func decodeAndValidate(r *http.Request, v validator.InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	// peform validation on the InputValidation implementation
	return v.Validate(r)
}