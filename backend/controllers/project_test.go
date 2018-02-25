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
)

func TestShowProject(t *testing.T)  {
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


func TestUpdateProject(t *testing.T)  {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	b := bytes.NewBufferString(`{"name" : "update"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/update/77fd6107-1889-11e8-8547-00224d6a96db/", b)
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
