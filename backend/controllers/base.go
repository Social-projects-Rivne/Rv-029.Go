package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"net/http"
	"log"
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

type Response struct {
	Status bool
	Message string
	StatusCode int
	Data interface{}
}

func (r *Response) Success(w http.ResponseWriter) {
	r.Status = true

	if r.StatusCode == 0 {
		r.StatusCode = http.StatusOK
	}

	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Printf("Error while json decode: %q", err.Error())
		r.StatusCode = http.StatusInternalServerError
		r.Message = "Horrible error"
		r.Failed(w)
	}

	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (r *Response) Failed(w http.ResponseWriter) {
	r.Status = false

	if r.StatusCode == 0 {
		r.StatusCode = http.StatusBadRequest
	}

	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Printf("Error while json decode: %q", err.Error())
	}

	w.WriteHeader(r.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
