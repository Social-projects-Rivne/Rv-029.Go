package validator

import "net/http"

type RegisterRequestData struct {
	FirstName string
	LastName string
	Email string
	Password string
	ConfirmPassword string
}

func (d *RegisterRequestData) Validate(r *http.Request) error {
	var err error
	err = ValidateEmail(d.Email)
	if err != nil {
		return err
	}

	//TODO:
	//err = ValidateEmailUnique(d.Email)
	//if err != nil {
	//	return err
	//}

	//TODO:
	//err = ValidatePasswordConfirmed(d.Password, d.ConfirmPassword)
	//if err != nil {
	//	return err
	//}

	err = ValidateMinLenght(d.FirstName, 3)
	if err != nil {
		return err
	}

	err = ValidateMinLenght(d.LastName, 3)
	if err != nil {
		return err
	}

	return nil
}