package controllers

import (
	"errors"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ######## DELETE BOARD ########

func TestDeleteSprintSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/sprint/delete/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteSprint)
	r.Handle("/project/board/sprint/delete/{sprint_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteSprintBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/sprint/delete/does-not-valid-id/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteSprint)
	r.Handle("/project/board/sprint/delete/{sprint_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestDeleteSprintDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)

	customError := errors.New("DB Error")
	mockSprintCRUD.EXPECT().Delete(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/sprint/delete/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteSprint)
	r.Handle("/project/board/sprint/delete/{sprint_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
