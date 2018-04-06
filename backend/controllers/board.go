package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/helpers"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func CreateBoard(w http.ResponseWriter, r *http.Request) {

	boardRequestData := new(validator.BoardCreateRequestData)
	err := decodeAndValidate(r, boardRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error(), StatusCode: http.StatusUnprocessableEntity}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/board/CreateBoard: %v", err)
		response := helpers.Response{Message: "Project ID is not valid"}
		response.Failed(w)
		return
	}

	project := models.Project{}
	project.UUID = projectId
	err = models.ProjectDB.FindByID(&project)
	if err != nil {
		response := helpers.Response{Message: "There is no Project with current ID"}
		response.Failed(w)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		project.UUID,
		project.Name,
		boardRequestData.Name,
		boardRequestData.Desc,
		nil,
		time.Now(),
		time.Now(),
	}

	err = models.BoardDB.Insert(&board)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has created"}
	response.Success(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {

	var boardRequestData validator.BoardUpdateRequestData
	err := decodeAndValidate(r, &boardRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/board/UpdateBoard: %v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = models.BoardDB.FindByID(&board)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	board.Name = boardRequestData.Name
	board.Desc = boardRequestData.Desc
	board.UpdatedAt = time.Now()
	err = models.BoardDB.Update(&board)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has updated"}
	response.Success(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/board/DeleteBoard: %v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = models.BoardDB.Delete(&board)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "Board has deleted"}
	response.Success(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/board/SelectBoard: %v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = id
	err = models.BoardDB.FindByID(&board)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Data: board}
	response.Success(w)
}

func BoardsList(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])
	if err != nil {
		log.Printf("Error in controllers/board/BoardList: %v", err)
		response := helpers.Response{Message: "Project ID is not valid"}
		response.Failed(w)
		return
	}

	boardsList, err := models.BoardDB.List(projectId)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Data: boardsList}
	response.Success(w)
}

func AssignUserToBoard(w http.ResponseWriter, r *http.Request) {

	var UserBoardRequestData validator.UserBoardRequestData
	err := decodeAndValidate(r, &UserBoardRequestData)
	if err != nil {
		response := helpers.Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {
		log.Printf("Error in controllers/board: %v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	userId, err := gocql.ParseUUID(UserBoardRequestData.UserId)
	if err != nil {
		log.Printf("Error in controllers/board: %v", err)
		response := helpers.Response{Message: "User ID is not valid"}
		response.Failed(w)
		return
	}

	err = models.BoardDB.AddUserToBoard(userId, UserBoardRequestData.Email, boardId)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "User successfully added to the board"}
	response.Success(w)

}

//DeleteUserFromBoard removes user from board
func DeleteUserFromBoard(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])
	if err != nil {	
		log.Printf("Error in controllers/board: %v", err)
		response := helpers.Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	userId, err := gocql.ParseUUID(vars["user_id"])
	if err != nil {
		log.Printf("Error in controllers/board: %v", err)
		response := helpers.Response{Message: "User ID is not valid"}
		response.Failed(w)
		return
	}

	err = models.BoardDB.DeleteUserFromBoard(userId, boardId)
	if err != nil {
		response := helpers.Response{Message: "Error while accessing to database", StatusCode: http.StatusInternalServerError}
		response.Failed(w)
		return
	}

	response := helpers.Response{Message: "User successfully delete from board"}
	response.Success(w)

}
