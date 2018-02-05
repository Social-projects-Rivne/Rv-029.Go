package validator

import (
	"testing"
	"net/http/httptest"
	"encoding/json"
	"strings"
)

//Validate ..
func TestConfirmRegistrationRequestData_Validate(t *testing.T) {
	input := &struct {
		Token string
	}{
		Token: "someStringToCheckIfExists",
	}
	jsonInput, _ := json.Marshal(input)

	request := httptest.NewRequest("GET", "http://localhost/", strings.NewReader(string(jsonInput)))

	validator := ConfirmRegistrationRequestData{}
	if err := json.NewDecoder(request.Body).Decode(validator); err != nil {
		t.Fatal(err.Error())
	}
	defer request.Body.Close()

	validationError := validator.Validate(request)
	if validationError != nil {
		t.Fatal(validationError.Error())
	}
}