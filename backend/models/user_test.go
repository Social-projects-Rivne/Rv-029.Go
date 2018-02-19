package models

import(
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
)

func TestInsert(t *testing.T){

mockCtrl := gomock.NewController(t)
defer mockCtrl.Finish()

mockUserer := mocks.NewMockUserer(mockCtrl)

mockUserer.EXPECT().Insert().Return().Times(1)

}