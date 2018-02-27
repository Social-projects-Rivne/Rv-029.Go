package controllers

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"net/http/httptest"
	"net/http"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gorilla/mux"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"encoding/json"
	"bytes"
	"errors"
)

// show project tests

func TestShowProjectSuccess(t *testing.T)  {
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
		t.Errorf("handler returned unexpected body: got %v want JSON",res.Body.String())
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


func TestShowProjectDBError(t *testing.T)  {
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

func TestUpdateProjectBadVariable(t *testing.T)  {


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

func TestUpdateProjectValidate(t *testing.T)  {


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

func TestUpdateProjectDBError(t *testing.T)  {
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


func TestUpdateProjectSuccess(t *testing.T)  {
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

func TestProjectsListSuccess(t *testing.T)  {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	//TODO return parameters
	mockProjectCRUD.EXPECT().GetProjectList(gomock.Any()).Return(nil ,nil).Times(1)


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


func TestProjectsListDBError(t *testing.T)  {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)

	customError := errors.New("DB Error")
	mockProjectCRUD.EXPECT().GetProjectList(gomock.Any()).Return(nil ,customError).Times(1)


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


func TestCreateProjectSuccess(t *testing.T)  {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	mockUserCRUD := mocks.NewMockUserCRUD(mockCtrl)
	models.InitUserDB(mockUserCRUD)
	mockUserCRUD.EXPECT().AddRoleToProject(gomock.Any(), gomock.Any() , gomock.Any()).Return(nil).Times(1)

	// TODO add context user to body
	boby := bytes.NewBufferString(`{"name" : "insert project"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/create/", boby)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateProject)
	r.Handle("/project/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)


	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}