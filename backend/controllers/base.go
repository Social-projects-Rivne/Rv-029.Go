package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"net/http"
	"log"
)

// decodeAndValidate - entry point for deserialization and validation of the submission input
func decodeAndValidate(r *http.Request, v validator.InputValidation) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	return v.Validate(r)
}

