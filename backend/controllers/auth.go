package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
)

type loginResponse struct {
	Status bool
	Message string
	Token string
}

func Login(w http.ResponseWriter, r *http.Request)  {
	var credentials struct{
		Email string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response := loginResponse{
			Status: false,
			Message: "Bad Request",
		}
		jsonResponse, _ := json.Marshal(response)

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

	fmt.Println("%v", credentials)
	user := models.User{}
	//TODO: find user by email

	if user.Password != password.EncodePassword(credentials.Password, user.Salt) {
		response := loginResponse{
			Status: false,
			Message: "There is no such user with email and password combination.",
		}
		jsonResponse, _ := json.Marshal(response)

		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}

	// generate jwt token from user claims
	token, _ := jwt.GenerateToken(&user)
	response := loginResponse{
		Status: true,
		Message: "You was successfully authenticated.",
		Token: token,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, _ := json.Marshal(response)

	w.Write(jsonResponse)


	//TODO: check if password is valid
	//TODO: generate jwt token
}
