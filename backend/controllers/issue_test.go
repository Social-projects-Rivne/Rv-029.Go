package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)


func TestDeleteIssueSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)
	mockIssueCRUD.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/issue/delete/9228322a-1ca2-33e8-ba28-c06e22a3322c/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteIssue)
	r.Handle("/project/board/issue/delete/{issue_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Issue has deleted","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestDeleteIssueBadVariable(t *testing.T) {

	// mockCtrl := gomock.NewController(t)
	// defer mockCtrl.Finish()

	// mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	// models.InitIssueDB(mockIssueCRUD)
	// mockIssueCRUD.EXPECT().Delete(gomock.Any()).Return(nil).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/issue/delete/does-not-valid-id/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteIssue)
	r.Handle("/project/board/issue/delete/{issue_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusUnprocessableEntity	 {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: invalid UUID \"does-not-valid-id\"","StatusCode":422,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestDeleteIssueDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)

	customError := errors.New("DB Error")
	mockIssueCRUD.EXPECT().Delete(gomock.Any()).Return(customError).Times(1)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "/project/board/issue/delete/9228322a-1ca2-33e8-ba28-c06e22a3322c/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(DeleteIssue)
	r.Handle("/project/board/issue/delete/{issue_id}/", handler).Methods("DELETE")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestCreateIssueDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	customError := errors.New("DB Error")
	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)	
	mockIssueCRUD.EXPECT().Insert(gomock.Any()).Return(customError).Times(1)

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/board/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/create/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(StoreIssue)
	r.Handle("/project/board/{board_id}/issue/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestCreateIssueBadVariable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/board/does-not-valid-id/issue/create/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(StoreIssue)
	r.Handle("/project/board/{board_id}/issue/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: invalid UUID \"does-not-valid-id\"","StatusCode":422,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestCreateIssueSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardCRUD := mocks.NewMockBoardCRUD(mockCtrl)
	models.InitBoardDB(mockBoardCRUD)
	mockBoardCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)
	mockIssueCRUD.EXPECT().Insert(gomock.Any()).Return(nil).Times(1)

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/project/board/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/create/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(StoreIssue)
	r.Handle("/project/board/{board_id}/issue/create/", handler).Methods("POST")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Issue has created","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestUpdateIssueSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)
	mockIssueCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockIssueCRUD.EXPECT().Update(gomock.Any()).Return(nil).Times(1)

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/board/issue/update/e0c72df7-16f3-11e8-8053-00224d6a96e3/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateIssue)
	r.Handle("/project/board/issue/update/{issue_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Issue has updated","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestUpdateIssueBadVariable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/board/issue/update/does-not-valid-id/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateIssue)
	r.Handle("/project/board/issue/update/{issue_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: invalid UUID \"does-not-valid-id\"","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestUpdateIssueFindByIDDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)	
	mockIssueCRUD.EXPECT().FindByID(gomock.Any()).Return(customError).Times(1)	


	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/board/issue/update/e0c72df7-16f3-11e8-8053-00224d6a96e3/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateIssue)
	r.Handle("/project/board/issue/update/{issue_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestUpdateIssueUpdateDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)	
	mockIssueCRUD.EXPECT().FindByID(gomock.Any()).Return(nil).Times(1)
	mockIssueCRUD.EXPECT().Update(gomock.Any()).Return(customError).Times(1)	


	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", "/project/board/issue/update/e0c72df7-16f3-11e8-8053-00224d6a96e3/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(UpdateIssue)
	r.Handle("/project/board/issue/update/{issue_id}/", handler).Methods("PUT")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestBoardIssueslistSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)
	mockIssueCRUD.EXPECT().GetBoardBacklogIssuesList(gomock.Any()).Return(nil,nil).Times(1)

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/list/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardIssueslist)
	r.Handle("/project/board/{board_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Done","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestSprintIssueslistSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)
	mockIssueCRUD.EXPECT().GetSprintIssueList(gomock.Any()).Return(nil,nil).Times(1)

	requestData := &struct {
		Name string
		Description string
	}{
		"issueName",
		"issueDescription",
	}

	body, _ := json.Marshal(requestData)

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/sprint/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/list/", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintIssueslist)
	r.Handle("/project/board/sprint/{sprint_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"Status":true,"Message":"Done","StatusCode":200,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestBoardIssueslistBadVariable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/does-not-valid-id/issue/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardIssueslist)
	r.Handle("/project/board/{board_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: invalid UUID \"does-not-valid-id\"","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestSprintIssueslistBadVariable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/sprint/does-not-valid-id/issue/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintIssueslist)
	r.Handle("/project/board/sprint/{sprint_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: invalid UUID \"does-not-valid-id\"","StatusCode":400,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}


func TestBoardIssueslistDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)	
	mockIssueCRUD.EXPECT().GetBoardBacklogIssuesList(gomock.Any()).Return(nil,customError).Times(1)	

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(BoardIssueslist)
	r.Handle("/project/board/{board_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}

func TestSprintIssueslistUpdateDBError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	customError := errors.New("DB Error")
	mockIssueCRUD := mocks.NewMockIssueCRUD(mockCtrl)
	models.InitIssueDB(mockIssueCRUD)	
	mockIssueCRUD.EXPECT().GetSprintIssueList(gomock.Any()).Return(nil,customError).Times(1)	

	r := *mux.NewRouter()
	res := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/project/board/sprint/93ab624a-1cb2-228a-ba34-c06ebf83322c/issue/list/", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(SprintIssueslist)
	r.Handle("/project/board/sprint/{sprint_id}/issue/list/", handler).Methods("GET")
	r.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"Status":false,"Message":"Error occured in controllers/issue.go error: DB Error","StatusCode":500,"Data":null}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}
