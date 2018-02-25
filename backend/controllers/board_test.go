package controllers

import (
	"encoding/json"
	"errors"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeleteBoardSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/delete/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteBoard)
	r.Handle("/project/board/delete/{board_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Board has deleted","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestDeleteBoardBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/delete/does-not-valid-id/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteBoard)
	r.Handle("/project/board/delete/{board_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Board ID is not valid","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestDeleteBoardDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)

	customError := errors.New("DB Error")
	mockBoardCRUD.EXPECT().Delete(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/delete/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteBoard)
	r.Handle("/project/board/delete/{board_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error while accessing to database","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestCreateBoardSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockInputValidator := mocks.NewMockInputValidation(mockCtrl)
	mockInputValidator.EXPECT().Validate(gomock.Any()).Return(nil).Times(1)

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	requestData := &struct {
		Name string
		Desc string
	}{
		"boardName",
		"boardDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/9325624a-0ba2-22e8-ba34-c06ebf83499a/board/create/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateBoard)
	r.Handle("/project/{project_id}/board/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Board has created","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}
