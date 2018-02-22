package controllers

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"net/http"
	"net/http/httptest"
)

func TestDeleteBoard(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCRUD := mocks.NewMockCRUD(mockCtrl)
	testBoard := NewModel{mockCRUD}
	mockCRUD.EXPECT().Delete().Return(nil).Times(1)
	testBoard.Delete()

	req, err := http.NewRequest("DELETE", "/project/board/delete/{9325624a-0ba2-22e8-ba34-c06ebf83499a}", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBoard)
	handler.ServeHTTP(res, req)

	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if res.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			res.Body.String(), expected)
	}
}