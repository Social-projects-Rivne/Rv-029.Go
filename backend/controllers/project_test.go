package controllers

import (
	"testing"
	"io"
	"github.com/golang/mock/gomock"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
)

var reader   io.Reader

func TestCreateProject(t *testing.T) {

	//server := httptest.NewServer(router.Router) //Creating new server with the user handlers
	//
	//projectUrl := fmt.Sprintf("%s/project/create", server.URL)
	//
	//projectJson := `{ "name" : "project number one" }`
	//
	//reader = strings.NewReader(projectJson) //Convert string to reader
	//
	//request, err := http.NewRequest("POST", projectUrl, reader)
	//
	//res, err := http.DefaultClient.Do(request)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockProjectStorage := mocks.NewMockBoardStorage(mockCtrl)

	project := &models.Project{}

	mockProjectStorage.EXPECT().Insert().Return(nil).Times(1)

	project.Insert()
	//if err != nil {
	//	t.Error(err) //Something is wrong while sending request
	//}
	//
	//if res.StatusCode != 201 {
	//	t.Errorf("Success expected: %d", res.StatusCode) //Uh-oh this means our test failed
	//}
}