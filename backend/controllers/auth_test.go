package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gocql/gocql"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/jwt"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/mail"
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
	config := &mail.SmtpMailerConfig{
		Connection: struct {
			Host     string
			Port     int
			Username string
			Password string
			Auth     string
			Tls      bool
		}{
			Host:     "smtp.mailtrap.io",
			Port:     465,
			Username: "7becbf096c9b34",
			Password: "deb0640e7fac43",
			Auth:     "plain",
			Tls:      true,
		},
		Sender: struct {
			Name  string
			Email string
		}{
			Name:  "Some Sender",
			Email: "sender@mail.com",
		},
	}

	mail.InitFromConfig(config)

	mail.Mailer = &mail.SmtpMailer{config}

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(nil).Times(1)
	mockUserCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	requestData := bytes.NewBufferString(`{"name": "Nigga", "surname": "Petrovich", "email": "assdf@gmail.com", "password": "zzz"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/register/", requestData)
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

func TestConfirmRegistrationSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
	id, _ := gocql.ParseUUID("566a70e6-1fac-11e8-b467-0ed5f89f718b")
	requestData := &struct {
		Token string
		UUID  gocql.UUID
	}{
		Token: "8934784566",
		UUID:  id,
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/confirm", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ConfirmRegistration)
	r.Handle("/auth/confirm", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":true,"Message":"Your account was successfully activated."}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestForgotPasswordSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	config := &mail.SmtpMailerConfig{
		Connection: struct {
			Host     string
			Port     int
			Username string
			Password string
			Auth     string
			Tls      bool
		}{
			Host:     "smtp.mailtrap.io",
			Port:     465,
			Username: "7becbf096c9b34",
			Password: "deb0640e7fac43",
			Auth:     "plain",
			Tls:      true,
		},
		Sender: struct {
			Name  string
			Email string
		}{
			Name:  "Some Sender",
			Email: "sender@mail.com",
		},
	}

	mail.InitFromConfig(config)

	mail.Mailer = &mail.SmtpMailer{config}

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(nil).Times(2)
	requestData := &struct {
		Email string
	}{
		Email: "nigga@gmail.com",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/forget-password", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ForgotPassword)
	r.Handle("/auth/forget-password", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestResetPasswordSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := models.User{
		Password: "8934784566",
		Salt:     "3SMtYvSg",
	}
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(nil).Times(1)
	mockUserCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)
	mockUserCRUD.EXPECT().CheckUserPassword(gomock.Any()).Return(user, nil).Times(1)

	requestData := &struct {
		Email    string
		Password string
		Token    string
	}{
		Email:    "nigga@gmail.com",
		Password: "8934784566",
		Token:    "8934784566",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/new-password", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ResetPassword)
	r.Handle("/auth/new-password", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetUserInfoSuccess(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	// userID, err := gocql.ParseUUID("550e8400-e29b-41d4-a716-446655440000")
	// user := models.User{UUID: userID}
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	//mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)

	models.InitUserDB(mockUserCRUD)
	// models.InitProjectDB(mockProjectCRUD)

	mockUserCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	// mockProjectCRUD.EXPECT().GetProjectsNamesList(gomock.Any()).Return(nil, nil).Times(1)
		

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/profile/9646324a-0aa2-11e8-ba15-b06ebf83499f", nil)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(GetUserInfo)
	r.Handle("/profile/{user_id}", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateUserInfoSuccess(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().UpdateFirstAndLastName(gomock.Any()).Return(nil).Times(1)
	body := bytes.NewBufferString(`{"name" : "name","surname" : "surname"}`)
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/profile/update", body)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(UpdateUserInfo)
	r.Handle("/profile/update", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":true,"Message":"Done","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestLoginDBError(t *testing.T) {
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

	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().CheckUserPassword(gomock.Any()).Return(user, customError).Times(1)

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

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestRegisterDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(customError).Times(1)
	mockUserCRUD.EXPECT().Insert(gomock.Any()).Return(customError).Times(1)

	requestData := bytes.NewBufferString(`{"name": "Nigga", "surname": "Petrovich", "email": "assdf@gmail.com", "password": "zzz"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/register/", requestData)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(Register)
	r.Handle("/auth/register/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestConfirmRegistrationDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().Update(gomock.Any()).Return(customError).Times(1)

	id, _ := gocql.ParseUUID("566a70e6-1fac-11e8-b467-0ed5f89f718b")
	requestData := &struct {
		Token string
		UUID  gocql.UUID
	}{
		Token: "8934784566",
		UUID:  id,
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/confirm", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ConfirmRegistration)
	r.Handle("/auth/confirm", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestForgotPasswordDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(customError).Times(2)

	requestData := &struct {
		Email string
	}{
		Email: "nigga@gmail.com",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/forget-password", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ForgotPassword)
	r.Handle("/auth/forget-password", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestResetPasswordDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := models.User{}
	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().FindByEmail(gomock.Any()).Return(customError).Times(1)
	mockUserCRUD.EXPECT().CheckUserPassword(gomock.Any()).Return(user, customError).Times(1)

	requestData := &struct {
		Email    string
		Password string
		Token    string
	}{
		Email:    "nigga@gmail.com",
		Password: "8934784566",
		Token:    "8934784566",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/new-password", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ResetPassword)
	r.Handle("/auth/new-password", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetUserInfoDBError(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)

	models.InitUserDB(mockUserCRUD)

	mockUserCRUD.EXPECT().FindByID(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/profile/598de923-30c8-11e8-b80e-c85b76da292c", nil)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(GetUserInfo)
	r.Handle("/profile/{user_id}", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestUpdateUserInfoDBError(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().UpdateFirstAndLastName(gomock.Any()).Return(customError).Times(1)
	body := bytes.NewBufferString(`{"name" : "name","surname" : "surname"}`)
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/profile/update", body)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(UpdateUserInfo)
	r.Handle("/profile/update", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestLoginBadVariable(t *testing.T) {

	requestData := &struct {
		Password string
	}{
		"owner@gmail.com",
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

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestConfirmRegistrationBadVariable(t *testing.T) {
	requestData := &struct {
		Email string
	}{
		Email: "8934784566",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/confirm", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ConfirmRegistration)
	r.Handle("/auth/confirm", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: value cannot be empty string","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestForgotPasswordBadVariable(t *testing.T) {
	requestData := &struct {
		BadVariable string
	}{
		BadVariable: "nigga@gmail.com",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/auth/forget-password", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ForgotPassword)
	r.Handle("/auth/forget-password", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: Invalid Email","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestUpdateUserInfoBadVariable(t *testing.T) {

	body := bytes.NewBufferString(`{"name":""}`)
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/profile/update", body)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(UpdateUserInfo)
	r.Handle("/profile/update", handler).Methods("POST")
	r.ServeHTTP(res, req)

	expected := `{"Status":false,"Message":"Error occured in controllers/auth.go error: while decoding json error: User.FirstName is empty\n","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
