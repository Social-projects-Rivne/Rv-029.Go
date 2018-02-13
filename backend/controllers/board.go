package controllers

import (
	"log"
	"encoding/json"
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
		response := baseResponse{false, err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	projectID, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := baseResponse{false, "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		projectID,
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	err = board.Insert()

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		log.Printf("Error while accessing to database: %v",err)
		return
	}

	response := baseResponse{true, "Board has created"}
	response.Success(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardUpdateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := baseResponse{false, err.Error()}
		response.Failed(w)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := baseResponse{false, "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	board.FindByID()

	if boardRequestData.Name != "" {
		board.Name = boardRequestData.Name
	}

	if boardRequestData.Desc != "" {
		board.Desc = boardRequestData.Desc
	}

	board.UpdatedAt = time.Now()
	err = board.Update()

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := baseResponse{true, "Board has updated"}
	response.Success(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := baseResponse{false, "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.Delete()

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		return
	}

	response := baseResponse{true, "Board has deleted"}
	response.Success(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := baseResponse{false, "Board ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}
	board.ID = id
	board.FindByID()

	jsonResponse, _ := json.Marshal(board)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func BoardsList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := baseResponse{false, "Project ID is not valid"}
		response.Failed(w)
		return
	}

	board := models.Board{}

	boardsList, err := board.List(projectId)

	if err != nil {
		response := baseResponse{false, "Error while accessing to database"}
		response.Failed(w)
		return
	}

	jsonResponse, _ := json.Marshal(boardsList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
