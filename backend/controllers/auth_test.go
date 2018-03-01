package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gocql/gocql"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/password"
	gojwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

func TestLoginSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	jwt.Config = &jwt.JWTConfig{
		Claims:      gojwt.MapClaims{"iss": "Test App"},
		Secret:      "string",
		Ttl:         60,
		Refresh_ttl: 1080,
	}

	id, err := gocql.ParseUUID("9646324a-0aa2-11e8-ba15-b06ebf83499f")
	if err != nil {
		t.Fatal(err.Error())
	}

	user := models.User{
		Password: "43f86c69b7c612fdc72b2e3562d42fbd6be012940c9cf88b2cd50621adc144cb",
		Salt:     "3SMtYvSg",
		UUID:     id,
	}

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().CheckUserPassword(gomock.Any()).Return(user, nil).Times(1)

	requestData := &struct {
		Email    string
		Password string
	}{
		"owner@gmail.com",
		password.EncodeMD5("qwerty1234"),
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/login/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Login)
	r.Handle("/auth/login/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestRegisterSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(nil).Times(1)

	requestData := &struct {
		FirstName string 
		LastName  string 
		Email     string 
		Password  string 
	}{
		"Niggasdsa",
		"Shitdsd",
		"owner@gmail.com",
		password.EncodeMD5("qwerty1234"),
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/register/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Register)
	r.Handle("/auth/register/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

