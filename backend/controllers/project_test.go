package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

// show project tests

func TestShowProjectSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/show/7b1ef3e9-13e0-11e8-ba83-b06ebf83499f/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ShowProject)
	r.Handle("/project/show/{project_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// TODO refactor handler error
	response := helpers.Response{}
	body := []byte(res.Body.String())

	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("handler returned unexpected body: got %v want JSON", res.Body.String())
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			response.StatusCode, http.StatusOK)
	}

}

func TestShowProjectBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/show/invalid-uuid/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ShowProject)
	r.Handle("/project/show/{project_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestShowProjectDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)

	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().FindByID(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/show/7b1ef3e9-13e0-11e8-ba83-b06ebf83499f/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ShowProject)
	r.Handle("/project/show/{project_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

// update project tests

func TestUpdateProjectBadVariable(t *testing.T) {

	body := bytes.NewBufferString(`{"name" : "update"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/update/invalid-uuid/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateProject)
	r.Handle("/project/update/{project_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateProjectValidate(t *testing.T) {

	body := bytes.NewBufferString(`{"name" : ""}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/update/7b1ef3e9-13e0-11e8-ba83-b06ebf83499f/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateProject)
	r.Handle("/project/update/{project_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestUpdateProjectDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)

	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().Update(gomock.Any()).Return(customError).Times(1)

	boby := bytes.NewBufferString(`{"name" : "update"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/update/77fd6107-1889-11e8-8547-00224d6a96db/", boby)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateProject)
	r.Handle("/project/update/{project_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestUpdateProjectSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	boby := bytes.NewBufferString(`{"name" : "update"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/update/77fd6107-1889-11e8-8547-00224d6a96db/", boby)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateProject)
	r.Handle("/project/update/{project_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

// project list test

func TestProjectsListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	//TODO return parameters
	mockProjectCRUD.EXPECT().GetProjectList(gomock.Any()).Return(nil, nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ProjectsList)
	r.Handle("/project/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestProjectsListDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)

	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().GetProjectList(gomock.Any()).Return(nil, customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ProjectsList)
	r.Handle("/project/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

// Create Project test

func TestCreateProjectSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	// TODO CHeck type each parameter in AddRoleToProject
	mockUserCRUD.EXPECT().AddRoleToProject(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).Times(1)

	boby := bytes.NewBufferString(`{"name" : "insert project"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/create/", boby)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(CreateProject)
	r.Handle("/project/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestCreateProjectDBInsertError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().Insert(gomock.Any()).Return(customError).Times(1)

	boby := bytes.NewBufferString(`{"name" : "insert project"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/create/", boby)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(CreateProject)
	r.Handle("/project/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestCreateProjectDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	customError := errors.New("DB Error")
	// TODO CHeck type each parameter in AddRoleToProject
	mockUserCRUD.EXPECT().AddRoleToProject(gomock.Any(), gomock.Any(), gomock.Any()).Return(customError).Times(1)

	boby := bytes.NewBufferString(`{"name" : "insert project"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/create/", boby)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(CreateProject)
	r.Handle("/project/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

// Delete Project test

func TestDeleteProjectSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().DeleteProject(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/delete/77fd6107-1889-11e8-8547-00224d6a96db/", nil)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(DeleteProject)
	r.Handle("/project/delete/{project_id}/", handler).Methods("Delete")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestDeleteProjectDBDeleteError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().DeleteProject(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().Delete(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/delete/77fd6107-1889-11e8-8547-00224d6a96db/", nil)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(DeleteProject)
	r.Handle("/project/delete/{project_id}/", handler).Methods("Delete")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestDeleteProjectDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().DeleteProject(gomock.Any(), gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/delete/77fd6107-1889-11e8-8547-00224d6a96db/", nil)
	if err != nil {
		t.Fatal(err)
	}

	user, err := helpers.InitFakeUser()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.WithValue(req.Context(), "user", user)

	req = req.WithContext(ctx)

	handler := http.HandlerFunc(DeleteProject)
	r.Handle("/project/delete/{project_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

func TestDeleteProjectBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/delete/invalid-uuid/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ShowProject)
	r.Handle("/project/delete/{project_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}


// ######## ProjectUsersList ########


func TestProjectUsersListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().GetProjectUsersList(gomock.Any()).Return(nil,nil).Times(1)


	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/7b1ef3e9-13e0-11e8-ba83-b06ebf83499f/users/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ProjectUsersList)
	r.Handle("/{project_id}/users/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestProjectUsersListBadVariable(t *testing.T) {

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/Wrong/users/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ProjectUsersList)
	r.Handle("/{project_id}/users/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

}

func TestProjectUsersListDBError(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().GetProjectUsersList(gomock.Any()).Return(nil,errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/7b1ef3e9-13e0-11e8-ba83-b06ebf83499f/users/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(ProjectUsersList)
	r.Handle("/{project_id}/users/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}