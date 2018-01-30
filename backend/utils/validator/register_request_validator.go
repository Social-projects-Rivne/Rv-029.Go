package validator

import "net/http"

type RegisterRequestData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (d *RegisterRequestData) Validate(r *http.Request) error {
	var err error
	err = ValidateEmail(d.Email)
	if err != nil {
		return err
	}

	err = ValidateEmailUnique(d.Email)
	if err != nil {
		return err
	}

	err = ValidateEmailExists(d.Email)
	if err != nil {
		return err
	}

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