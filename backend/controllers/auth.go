package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"log"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/mail"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
)

type errorResponse struct {
	Status  bool
	Message string
}

type loginResponse struct {
	Status  bool
	Message string
	Token   string
}

type registerResponse struct {
	Status  bool
	Message string
	User    models.User
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequestData validator.LoginRequestData

	err := decodeAndValidate(r, &loginRequestData)
	if err != nil {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user := models.User{}
	user.Email = loginRequestData.Email
	if err := models.UserDB.FindByEmail(&user);err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	if user.Password != password.EncodePassword(loginRequestData.Password, user.Salt) {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: "There is no such user with email and password combination.",
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	// generate jwt token from user claims
	token, err := jwt.GenerateToken(&user)
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, err := json.Marshal(loginResponse{
		Status:  true,
		Message: "You was successfully authenticated.",
		Token:   token,
	})
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.Write(jsonResponse)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var registerRequestData validator.RegisterRequestData

	err := decodeAndValidate(r, &registerRequestData)
	if err != nil {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID:      gocql.TimeUUID(),
		Email:     registerRequestData.Email,
		FirstName: registerRequestData.FirstName,
		LastName:  registerRequestData.LastName,
		Salt:      salt,
		Password:  password.EncodePassword(registerRequestData.Password, salt),
		Role:      models.ROLE_USER,
		Status:    0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := models.UserDB.Insert(&user);err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	message := fmt.Sprintf("Hello %s,\nYou was successfully registered in \"Task manager\".\n To activate your account please go to the <a href=\"http://localhost/authorization/login/?token=%s&uuid=%s\">LINK</a>\n Your ID: %s\n Regards\n", user.FirstName, user.Password, user.UUID, user.UUID)
	mail.Mailer.Send(user.Email, user.FirstName, "Successfully Registered", message)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "You was successfully registered",
		User:    user,
	})
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return
}

func ConfirmRegistration(w http.ResponseWriter, r *http.Request) {
	var confirmRegistrationRequestData validator.ConfirmRegistrationRequestData

	err := decodeAndValidate(r, &confirmRegistrationRequestData)
	if err != nil {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}


		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	fmt.Println(confirmRegistrationRequestData)
	user := models.User{}
	user.Status = 1
	models.UserDB.Update(&user)

	jsonResponse, err := json.Marshal(struct {
		Status  bool
		Message string
	}{
		Status:  true,
		Message: "Your account was successfully activated.",
	})
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return
}

//ForgotPassword ..
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var forgotRequestData validator.ForgotPasswordRequestData

	err := decodeAndValidate(r, &forgotRequestData)
	if err != nil {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user := models.User{}
	user.Email = forgotRequestData.Email
	// TODO err handler
	err = models.UserDB.FindByEmail(&user)

	message := fmt.Sprintf("Hello %s,\nIt is your link to restore password <a href=\"http://localhost/authorization/new-password/%s\">LINK</a>\n", user.FirstName, user.Password)
	mail.Mailer.Send(user.Email, user.FirstName, "Successfully Registered", message)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "Your link sent",
	})
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return

}

//ResetPassword ..
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var resetRequestData validator.ResetPasswordRequestData
	fmt.Println(resetRequestData)
	err := decodeAndValidate(r, &resetRequestData)

	if err != nil {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: err.Error(),
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	user := models.User{}
	user.Email = resetRequestData.Email
	err = models.UserDB.FindByEmail(&user)

	// FIXME doesn't work properly
	if user.Password != resetRequestData.Token {
		jsonResponse, err := json.Marshal(errorResponse{
			Status:  false,
			Message: "Invalid reset token",
		})
		if err != nil{
			log.Printf("Error in controllers/auth error: %+v",err)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user.Password = password.EncodePassword(resetRequestData.Password, user.Salt)

	models.UserDB.Update(&user)

	jsonResponse, err := json.Marshal(registerResponse{
		Status:  true,
		Message: "Your password restored",
	})
	if err != nil{
		log.Printf("Error in controllers/auth error: %+v",err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return

}
