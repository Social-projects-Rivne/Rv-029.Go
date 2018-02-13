package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"net/http"
)

// decodeAndValidate - entry point for deserialization and validation
// of the submission input
func decodeAndValidate(r *http.Request, v validator.InputValidation) error {
	// json decode the payload - obviously this could be abstracted
	// to handle many content types
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}
	defer r.Body.Close()
	// perform validation on the InputValidation implementation
	return v.Validate(r)
}

type baseResponse struct {
	Status  bool
	Message string
}

func (b *baseResponse) Success(w http.ResponseWriter) {
	jsonResponse, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (b *baseResponse) Failed(w http.ResponseWriter) {
	jsonResponse, _ := json.Marshal(b)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
