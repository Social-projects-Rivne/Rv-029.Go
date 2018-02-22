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

type failedResponse struct {
	Status  bool
	Message string
}

func (b *failedResponse) send(w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(b)
	if err != nil{
		log.Printf("Error in controllers/base error: %+v",err)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

type successResponse struct {
	Status  bool
	Message string
	Data    interface{}
}

func (b *successResponse) send(w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(b)
	if err != nil{
		log.Printf("Error in controllers/base error: %+v",err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
