package validator

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type UpdateUserRequestData struct {
	*baseValidator
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Email     string `json:"email"`
}

func (d *UpdateUserRequestData) Validate(r *http.Request) error {
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
	if d.FirstName == "" {
		log.Println("while decoding json error: User.FirstName is empty")
		return errors.New(fmt.Sprintln("while decoding json error: User.FirstName is empty"))
	}
	if d.LastName == "" {
		log.Println("while decoding json error: User.LastName is empty")
		return errors.New(fmt.Sprintln("while decoding json error: User.LastName is empty"))
	}
	if d.Email == "" {
		log.Println("while decoding json error: User.Email is empty")
		return errors.New(fmt.Sprintln("while decoding json error: User.Email is empty"))
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
