package controllers

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/mocks"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockBoardStorage := mocks.NewMockBoardStorage(mockCtrl)

	board := &models.Board{}

	mockBoardStorage.EXPECT().Insert().Return(nil).Times(1)

	board.Insert()
}