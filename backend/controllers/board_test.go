package controllers

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

// ######## DELETE BOARD ########

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

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error while accessing to database","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

// ######## CREATE BOARD ########

func TestCreateBoardSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	
	models.InitProjectDB(mockProjectCRUD)
	models.InitBoardDB(mockBoardCRUD)
	
	mockProjectCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockBoardCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	body := bytes.NewBufferString(`{"name": "boardName", "description": "boardDescription"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/9325624a-0ba2-22e8-ba34-c06ebf83499a/board/create/", body)
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

func TestCreateBoardBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"name": "boardName", "description": "boardDescription"}`)

	req, err := http.NewRequest("POST", "/project/does-not-valid-id/board/create/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateBoard)
	r.Handle("/project/{project_id}/board/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateBoardDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectCRUD := mocks.NewMockProjectCRUD(mockCtrl)
	models.InitProjectDB(mockProjectCRUD)
	mockProjectCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().Insert(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"name": "boardName", "description": "boardDescription"}`)

	req, err := http.NewRequest("POST", "/project/9325624a-0ba2-22e8-ba34-c06ebf83499a/board/create/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateBoard)
	r.Handle("/project/{project_id}/board/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// ######## UPDATE BOARD ########

func TestUpdateBoardSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockBoardCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	body := bytes.NewBufferString(`{"name": "boardName", "desc": "boardDescription"}`)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/board/update/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateBoard)
	r.Handle("/project/board/update/{board_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateBoardBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"name": "boardName", "desc": "boardDescription"}`)

	req, err := http.NewRequest("PUT", "/project/board/update/does-not-valid-id/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateBoard)
	r.Handle("/project/board/update/{board_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestUpdateBoardDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockBoardCRUD.EXPECT().Update(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"name": "boardName", "desc": "boardDescription"}`)

	req, err := http.NewRequest("PUT", "/project/board/update/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateBoard)
	r.Handle("/project/board/update/{board_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// ######## SELECT BOARD ########

func TestSelectBoardSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/select/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectBoard)
	r.Handle("/project/board/select/{board_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSelectBoardBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/select/does-not-valid-id/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectBoard)
	r.Handle("/project/board/select/{board_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSelectBoardDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/select/9325624a-0ba2-22e8-ba34-c06ebf83499a/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectBoard)
	r.Handle("/project/board/select/{board_id}/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// ######## BOARD LIST ########

func TestBoardsListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().List(gomock.Any()).Return(nil, nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/9325624a-0ba2-22e8-ba34-c06ebf83499a/board/list/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardsList)
	r.Handle("/project/{project_id}/board/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestBoardsListBadVariable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/does-not-valid-id/board/list/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardsList)
	r.Handle("/project/{project_id}/board/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestBoardsListDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().List(gomock.Any()).Return(nil, errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/9325624a-0ba2-22e8-ba34-c06ebf83499a/board/list/", nil)

	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardsList)
	r.Handle("/project/{project_id}/board/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
