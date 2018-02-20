package controllers

import (
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardCreateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	project := models.Project{}
	project.UUID = projectId
	err = project.FindByID()

	if err != nil {
		response := Response{Message: "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		project.UUID,
		project.Name,
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	err = board.Insert()

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := Response{Message: "Board has created"}
	response.Success(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardUpdateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := Response{Message: err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.FindByID()

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	board.Name = boardRequestData.Name
	board.Desc = boardRequestData.Desc
	board.UpdatedAt = time.Now()
	err = board.Update()

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := Response{Message: "Board has updated"}
	response.Success(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := Response{Message: "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.Delete()

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := Response{Message: "Board has deleted"}
	response.Success(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := Response{Message: "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = id
	err = board.FindByID()

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := Response{Data: board}
	response.Success(w)
}

func BoardsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := Response{Message: "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}

	boardsList, err := board.List(projectId)

	if err != nil {
		response := Response{Message: "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := Response{Data: boardsList}
	response.Success(w)
}
