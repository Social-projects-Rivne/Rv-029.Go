package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

// ######## Delete Sprint ########

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

// ######## Create Sprint ########

func TestCreateSprintSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("POST", "/project/board/9325624a-0ba2-22e8-ba34-c06ebf83499a/sprint/create/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateSprint)
	r.Handle("/project/board/{board_id}/sprint/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestCreateSprintBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("POST", "/project/board/does-not-valid-id/sprint/create/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateSprint)
	r.Handle("/project/board/{board_id}/sprint/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateSprintDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().Insert(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("POST", "/project/board/9325624a-0ba2-22e8-ba34-c06ebf83499a/sprint/create/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(CreateSprint)
	r.Handle("/project/board/{board_id}/sprint/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

}

// ######## Update Sprint ########

func TestUpdateSprintSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockSprintCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("PUT", "/project/board/sprint/update/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateSprint)
	r.Handle("/project/board/sprint/update/{sprint_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdateSprintBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("POST", "/project/board/sprint/update/does-not-valid-id/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateSprint)
	r.Handle("/project/board/sprint/update/{sprint_id}/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestUpdateSprintDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockSprintCRUD.EXPECT().Update(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("PUT", "/project/board/sprint/update/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateSprint)
	r.Handle("/project/board/sprint/update/{sprint_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// ######## Select Sprint ########

func TestSelectSprintSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/sprint/show/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectSprint)
	r.Handle("/project/board/sprint/show/{sprint_id}/", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSelectSprintBadVariable(t *testing.T) {

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/sprint/show/does-not-valid-id/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectSprint)
	r.Handle("/project/board/sprint/show/{sprint_id}/", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSelectSprintDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().FindByID(gomock.Any()).Return(errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/sprint/show/9325624a-0ba2-22e8-ba34-c06ebf83499a/", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SelectSprint)
	r.Handle("/project/board/sprint/show/{sprint_id}/", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// ######## Sprints List ########

func TestSprintsListSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	sprints := []models.Sprint{
		{
			ID:          gocql.TimeUUID(),
			ProjectId:   gocql.TimeUUID(),
			BoardId:     gocql.TimeUUID(),
			ProjectName: "testProjectName",
			BoardName:   "testBoardName",
			Goal:        "testGoal",
			Desc:        "testDescription",
			Status:      "testStatus",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().List(gomock.Any()).Return(sprints, nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/9325624a-0ba2-22e8-ba34-c06ebf83499a/sprint/list", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintsList)
	r.Handle("/project/board/{board_id}/sprint/list", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	sprintsOut := make([]models.Sprint, 0)
	json.Unmarshal(res.Body.Bytes(), &sprintsOut)

	if reflect.DeepEqual(sprints, sprintsOut) {
		t.Errorf("handler returned wrong sprints list: got %v want %v",
			sprintsOut, sprints)
	}

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSprintsListBadVariable(t *testing.T) {
	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/does-not-valid-id/sprint/list", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintsList)
	r.Handle("/project/board/{board_id}/sprint/list", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSprintsListDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockSprintCRUD := mocks.NewMockSprintCRUD(mockCtrl)
	models.InitSprintDB(mockSprintCRUD)
	mockSprintCRUD.EXPECT().List(gomock.Any()).Return(nil, errors.New("DB Error")).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()

	body := bytes.NewBufferString(`{"goal": "sprintName", "desc": "SprintDescription"}`)

	req, err := http.NewRequest("GET", "/project/board/9325624a-0ba2-22e8-ba34-c06ebf83499a/sprint/list", body)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintsList)
	r.Handle("/project/board/{board_id}/sprint/list", handler).Methods("GEt")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}
