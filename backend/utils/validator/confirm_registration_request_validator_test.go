package validator

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
	//"fmt"
)

//Validate ..
func TestConfirmRegistrationRequestData_Validate_Success(t *testing.T) {
	input := &struct {
		Token string
	}{
		Token: "someStringToCheckIfExists",
	}
	jsonInput, _ := json.Marshal(input)

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
	jsonInput, _ := json.Marshal(input)

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
