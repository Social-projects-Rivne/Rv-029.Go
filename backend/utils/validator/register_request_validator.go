package validator

import (
	"fmt"
	"log"
	"net/http"
)

type RegisterRequestData struct {
	*baseValidator
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (d *RegisterRequestData) Validate(r *http.Request) error {
	var err error
	err = d.ValidateEmail(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)
		return err
	}

	err = d.ValidateEmailUnique(d.Email)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)
		
		return err
	}

	// err = d.ValidateEmailExists(d.Email)
	// if err != nil {
	// 	return err
	// }

	//TODO:
	//err = d.ValidatePasswordConfirmed(d.Password, d.ConfirmPassword)
	//if err != nil {
	//	return err
	//}

	err = d.ValidateMinLenght(d.FirstName, 3)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)		
		return err
	}

	err = d.ValidateMinLenght(d.LastName, 3)
	if err != nil {
		log.Printf("Error in utils/validator/register_request_validator.go error: %+v", err)		
		return err
	}

	return nil
}
