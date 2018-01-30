package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/mail"
	"github.com/gocql/gocql"
	"time"
	"fmt"
)

type errorResponse struct {
	Status bool
	Message string
}

type loginResponse struct {
	Status bool
	Message string
	Token string
}

type registerResponse struct {
	Status bool
	Message string
	User models.User
}

func Login(w http.ResponseWriter, r *http.Request)  {
	var loginRequestData validator.LoginRequestData

	err := decodeAndValidate(r, &loginRequestData)
	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			Status: false,
			Message: err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	user := models.User{}
	user.FindByEmail(loginRequestData.Email)

	if user.Password != password.EncodePassword(loginRequestData.Password, user.Salt) {
		jsonResponse, _ := json.Marshal(errorResponse{
			Status: false,
			Message: "There is no such user with email and password combination.",
		})

		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	// generate jwt token from user claims
	token, _ := jwt.GenerateToken(&user)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(loginResponse{
		Status: true,
		Message: "You was successfully authenticated.",
		Token: token,
	})

	w.Write(jsonResponse)
}

func Register(w http.ResponseWriter, r *http.Request)  {
	var registerRequestData validator.RegisterRequestData

	err := decodeAndValidate(r, &registerRequestData)
	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			Status: false,
			Message: err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}

	salt := password.GenerateSalt(8)
	user := models.User{
		UUID: gocql.TimeUUID(),
		Email: registerRequestData.Email,
		FirstName: registerRequestData.FirstName,
		LastName: registerRequestData.LastName,
		Salt: salt,
		Password: password.EncodePassword(registerRequestData.Password, salt),
		Role: models.ROLE_USER,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user.Insert()

	message := fmt.Sprintf("Hello %s,\nYou was successfully registered in \"Task manager\"\n. Your ID: %s\n Regards\n", user.FirstName, user.UUID)
	mail.Mailer.Send(user.Email, user.FirstName, "Successfully Registered", message)

	jsonResponse, _ := json.Marshal(registerResponse{
		Status: true,
		Message: "You was successfully registered",
		User: user,
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	return
}


func ForgotPassword(w http.ResponseWriter, r *http.Request){

	var forgotRequestData validator.ForgotPasswordRequestData

	err := decodeAndValidate(r, &forgotRequestData)
	if err != nil {
		jsonResponse, _ := json.Marshal(errorResponse{
			Status: false,
			Message: err.Error(),
		})

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
		return
	}
	user := &models.User{}

	user.FindByEmail(forgotRequestData.Email)

	

}