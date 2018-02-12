package controllers

import (
	"encoding/json"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/models"
	"github.com/Social-projects-Rivne/Rv-029.Go/backend/utils/validator"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type boardResponse struct {
	Status  bool
	Message string
}

func (b *boardResponse) sendSuccess(w http.ResponseWriter) {
	jsonResponse, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (b *boardResponse) sendFailed(w http.ResponseWriter, err error) {
	log.Printf("%v: %q", b.Message, err.Error())

	jsonResponse, _ := json.Marshal(b)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func CreateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardCreateRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := boardResponse{false, err.Error()}
		response.sendFailed(w, err)
		return
	}

	vars := mux.Vars(r)
	projectId, err := gocql.ParseUUID(vars["project_id"])

	if err != nil {
		response := boardResponse{false, "Project ID is not valid"}
		response.sendFailed(w, err)
		return
	}

	board := models.Board{
		gocql.TimeUUID(),
		projectId,
		boardRequestData.Name,
		boardRequestData.Desc,
		time.Now(),
		time.Now(),
	}

	err = board.Insert()

	if err != nil {
		response := boardResponse{false, "Error while accessing to database"}
		response.sendFailed(w, err)
		return
	}

	response := boardResponse{true, "Board has created"}
	response.sendSuccess(w)
}

func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	var boardRequestData validator.BoardBaseRequestData

	err := decodeAndValidate(r, &boardRequestData)

	if err != nil {
		response := boardResponse{false, err.Error()}
		response.sendFailed(w, err)
		return
	}

	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := boardResponse{false, "Project ID is not valid"}
		response.sendFailed(w, err)
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
		response := boardResponse{false, "Error while accessing to database"}
		response.sendFailed(w, err)
		return
	}

	response := boardResponse{true, "Board has updated"}
	response.sendSuccess(w)
}

func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	boardId, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := boardResponse{false, "Project ID is not valid"}
		response.sendFailed(w, err)
		return
	}

	board := models.Board{}
	board.ID = boardId
	err = board.Delete()

	if err != nil {
		response := boardResponse{false, "Error while accessing to database"}
		response.sendFailed(w, err)
		return
	}

	response := boardResponse{true, "Board has deleted"}
	response.sendSuccess(w)
}

func SelectBoard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := gocql.ParseUUID(vars["board_id"])

	if err != nil {
		response := boardResponse{false, "Board ID is not valid"}
		response.sendFailed(w, err)
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
		response := boardResponse{false, "Project ID is not valid"}
		response.sendFailed(w, err)
		return
	}

	board := models.Board{}

	boardsList, err := board.List(projectId)

	if err != nil {
		response := boardResponse{false, "Error while accessing to database"}
		response.sendFailed(w, err)
		return
	}

	jsonResponse, _ := json.Marshal(boardsList)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
