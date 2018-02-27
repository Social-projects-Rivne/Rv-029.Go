package validator

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	"log"
	//"fmt"
)

//Validate ..
func TestConfirmRegistrationRequestData_Validate_Success(t *testing.T) {
	input := &struct {
		Token string
	}{
		Token: "someStringToCheckIfExists",
	}
	jsonInput, err := json.Marshal(input)
	if err != nil{
		log.Printf("Error in utils/validator/board_update_request_validator_test.go error: %+v",err)
	}

	request := httptest.NewRequest("GET", "http://localhost/", strings.NewReader(string(jsonInput)))

	validator := ConfirmRegistrationRequestData{}
	if err := json.NewDecoder(request.Body).Decode(&validator); err != nil {
		t.Fatal(err.Error())
	}
	defer request.Body.Close()

	validationError := validator.Validate(request)
	if validationError != nil {
		t.Fatal(validationError.Error())
	}
}

//Validate ..
func TestConfirmRegistrationRequestData_Validate_Error(t *testing.T) {
	input := &struct {
		Token string
	}{
	//Token: nil,
	}
	jsonInput, err := json.Marshal(input)
	if err != nil{
		log.Printf("Error in utils/validator/confirm_registration_request_validator_test.go error: %+v",err)
	}

	request := httptest.NewRequest("GET", "http://localhost/", strings.NewReader(string(jsonInput)))

	validator := ConfirmRegistrationRequestData{}
	if err := json.NewDecoder(request.Body).Decode(&validator); err != nil {
		t.Fatal(err.Error())
	}
	defer request.Body.Close()

	validationError := validator.Validate(request)
	if validationError == nil {
		t.Fatal(errors.New("validator passed nil value"))
	}
}
